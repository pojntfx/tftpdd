package workers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pin/tftp"
)

// TFTPD is a TFTP server.
type TFTPD struct {
	bindAddress, root string
	instance          *tftp.Server
}

// NewTFTPD creates a new TFTP server.
func NewTFTPD(bindAddress, root string) *TFTPD {
	return &TFTPD{
		bindAddress: bindAddress,
		root:        root,
	}
}

func (t *TFTPD) readHandler(filename string, rf io.ReaderFrom) error {
	file, err := os.Open(filepath.Join(t.root, filename))
	if err != nil {
		return err
	}

	if _, err := rf.ReadFrom(file); err != nil {
		return err
	}

	return nil
}

// Not implemented
func (t *TFTPD) writeHandler(filename string, wt io.WriterTo) error {
	return nil
}

// Start starts the TFTP server and blocks until `Stop` is called
func (t *TFTPD) Start() error {
	t.instance = tftp.NewServer(t.readHandler, t.writeHandler)
	return t.instance.ListenAndServe(t.bindAddress)
}

// Stop stops the TFTP server
func (t *TFTPD) Stop() {
	t.instance.Shutdown()
}
