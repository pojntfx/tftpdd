package svc

//go:generate mkdir -p ../proto/generated
//go:generate sh -c "protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	TFTPDD "github.com/pojntfx/tftpdd/pkg/proto/generated"
	"github.com/pojntfx/tftpdd/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TFTPDManager manages TFTP servers.
type TFTPDManager struct {
	TFTPDD.UnimplementedTFTPDDManagerServer
	baseDir string
	workers map[string]*workers.TFTPD
}

// NewTFTPDManager creates a new TFTPDManager.
func NewTFTPDManager(baseDir string) *TFTPDManager {
	return &TFTPDManager{
		baseDir: baseDir,
		workers: make(map[string]*workers.TFTPD),
	}
}

// Stop stops the workers the TFTPDManager manages.
func (t *TFTPDManager) Stop() {
	for _, worker := range t.workers {
		worker.Stop()
	}
}

// Create creates a TFTP server.
func (t *TFTPDManager) Create(ctx context.Context, req *TFTPDD.TFTPD) (*TFTPDD.TFTPDId, error) {
	log.Info("Starting TFTP server")

	device, err := net.InterfaceByName(req.GetDevice())
	if err != nil {
		return nil, status.Error(codes.NotFound, "device not found")
	}

	addrs, err := device.Addrs()
	if err != nil {
		return nil, err
	}

	bindAddress := fmt.Sprintf("%v:%v", strings.Split(addrs[0].String(), "/")[0], req.GetPort())

	id := uuid.NewV4().String()

	dir := filepath.Join(t.baseDir, id)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return nil, err
	}

	worker := workers.NewTFTPD(bindAddress, dir)

	go func() {
		if err := worker.Start(); err != nil {
			log.Error("Error while starting TFTP server", rz.Err(err))
		}
	}()

	t.workers[id] = worker

	return &TFTPDD.TFTPDId{
		Id: id,
	}, nil
}

// List lists managed TFTP servers.
func (t *TFTPDManager) List(ctx context.Context, req *TFTPDD.TFTPDManagerListArgs) (*TFTPDD.TFTPDManagerListReply, error) {
	log.Info("Listing TFTP servers")

	var res []*TFTPDD.TFTPDManaged
	for id, worker := range t.workers {
		outWorker := &TFTPDD.TFTPDManaged{
			Id:            id,
			ListenAddress: worker.GetBindAddress(),
		}

		res = append(res, outWorker)
	}

	return &TFTPDD.TFTPDManagerListReply{
		TFTPDs: res,
	}, nil
}

// Get gets one of the managed TFTP servers.
func (t *TFTPDManager) Get(ctx context.Context, req *TFTPDD.TFTPDId) (*TFTPDD.TFTPDManaged, error) {
	log.Info("Getting TFTP server")

	var res *TFTPDD.TFTPDManaged
	for id, worker := range t.workers {
		if id == req.GetId() {
			res = &TFTPDD.TFTPDManaged{
				Id:            id,
				ListenAddress: worker.GetBindAddress(),
			}

			break
		}
	}

	if res != nil {
		return res, nil
	}

	msg := "TFTP server not found"

	log.Error(msg)

	return nil, status.Error(codes.NotFound, msg)
}
