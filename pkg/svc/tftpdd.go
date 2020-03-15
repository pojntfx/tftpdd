package svc

//go:generate mkdir -p ../proto/generated
//go:generate sh -c "protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	TFTPDD "github.com/pojntfx/tftpdd/pkg/proto/generated"
	"github.com/pojntfx/tftpdd/pkg/workers"
)

// TFTPDManager manages TFTP servers.
type TFTPDManager struct {
	TFTPDD.UnimplementedTFTPDDManagerServer
	workers map[string]*workers.TFTPD
}

// NewTFTPDManager creates a new TFTPDManager.
func NewTFTPDManager() *TFTPDManager {
	return &TFTPDManager{
		workers: make(map[string]*workers.TFTPD),
	}
}

// Stop stops the workers the TFTPDManager manages.
func (t *TFTPDManager) Stop() {
	for _, worker := range t.workers {
		worker.Stop()
	}
}
