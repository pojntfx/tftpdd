package workers

// TFTPD is a TFTP server.
type TFTPD struct {
	bindAddress, root string
}

// NewTFTPD creates a new TFTP server.
func NewTFTPD(bindAddress, root string) *TFTPD {
	return &TFTPD{
		bindAddress: bindAddress,
		root:        root,
	}
}
