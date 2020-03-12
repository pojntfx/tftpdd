package workers

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	listenAddrIn = "localhost:0"
)

var (
	tftpRoot = filepath.Join(os.TempDir(), "tftp")
)

func TestCreateTFTPD(t *testing.T) {
	d := NewTFTPD(listenAddrIn, tftpRoot)

	if d == nil {
		t.Error("New TFTPD is nil")
	}

	if d.bindAddress != listenAddrIn {
		t.Error("bindAddress not set correctly")
	}

	if d.root != tftpRoot {
		t.Error("root not set correctly")
	}
}
