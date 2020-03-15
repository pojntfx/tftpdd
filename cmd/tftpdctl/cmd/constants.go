package cmd

const (
	keyPrefix          = "tftpd."
	configFileDefault  = ""
	serverHostPortKey  = keyPrefix + "serverHostPort"
	configFileKey      = keyPrefix + "configFile"
	deviceKey          = keyPrefix + "device"
	portKey            = keyPrefix + "port"
	biosFilenameURLKey = keyPrefix + "biosFilenameURL"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
