package cmd

const (
	keyPrefix         = "tftpd."
	configFileDefault = ""
	serverHostPortKey = keyPrefix + "serverHostPort"
	configFileKey     = keyPrefix + "configFile"
	deviceKey         = keyPrefix + "device"
	portKey           = keyPrefix + "port"
	pxepackageURLKey  = keyPrefix + "pxepackageURL"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
