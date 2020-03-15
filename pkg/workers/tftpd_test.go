package workers

import (
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/pin/tftp"
)

const (
	testFileName     = "test.txt"
	testFileContents = "Hello, world!"
)

var (
	tftpRoot = filepath.Join(os.TempDir(), "tftp")
)

func getListenAddr() (string, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", err
	}

	if err := listener.Close(); err != nil {
		return "", err
	}

	return listener.Addr().String(), nil
}

func initTest() (string, error) {
	if err := os.RemoveAll(tftpRoot); err != nil {
		return "", err
	}

	if err := os.MkdirAll(tftpRoot, 0777); err != nil {
		return "", err
	}

	return getListenAddr()
}

func TestCreateTFTPD(t *testing.T) {
	listenAddr, err := initTest()
	if err != nil {
		t.Error(err)
	}

	d := NewTFTPD(listenAddr, tftpRoot)

	if d == nil {
		t.Error("New TFTPD is nil")
	}

	if d.bindAddress != listenAddr {
		t.Error("bindAddress not set correctly")
	}

	if d.root != tftpRoot {
		t.Error("root not set correctly")
	}
}

func TestStart(t *testing.T) {
	listenAddr, err := initTest()
	if err != nil {
		t.Error(err)
	}

	d := NewTFTPD(listenAddr, tftpRoot)

	go func() {
		if err := d.Start(); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(time.Millisecond * 100)

	if _, err := net.DialTimeout("udp", listenAddr, time.Second*1); err != nil {
		t.Error("Couldn't connect to TFTP server!", err)
	}

	d.Stop()
}

func TestGetFile(t *testing.T) {
	listenAddr, err := initTest()
	if err != nil {
		t.Error(err)
	}

	d := NewTFTPD(listenAddr, tftpRoot)

	go func() {
		if err := d.Start(); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(time.Millisecond * 100)

	if err := ioutil.WriteFile(filepath.Join(tftpRoot, testFileName), []byte(testFileContents), 0777); err != nil {
		t.Error(err)
	}

	c, err := tftp.NewClient(listenAddr)
	if err != nil {
		t.Error(err)
	}

	writer, err := c.Receive(testFileName, "octet")
	if err != nil {
		t.Error(err)
	}

	outPath := filepath.Join(os.TempDir(), testFileName)
	file, err := os.Create(outPath)
	if err != nil {
		t.Error(err)
	}

	if _, err := writer.WriteTo(file); err != nil {
		t.Error(err)
	}

	receivedContent, err := ioutil.ReadFile(outPath)
	if err != nil {
		t.Error(err)
	}

	if string(receivedContent) != testFileContents {
		t.Errorf("Received content %v does not match expected content %v", string(receivedContent), testFileContents)
	}

	d.Stop()
}

func TestGetBindAddress(t *testing.T) {
	listenAddr, err := initTest()
	if err != nil {
		t.Error(err)
	}

	d := NewTFTPD(listenAddr, tftpRoot)

	if d.GetBindAddress() != listenAddr {
		t.Errorf("GetBindAddress did not return the expected value (expected %v, got %v)", listenAddr, d.GetBindAddress())
	}
}
