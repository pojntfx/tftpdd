package main

import (
	"fmt"
	"strings"
	"time"

	constants "github.com/pojntfx/tftpdd/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
)

const (
	keyPrefix         = "tftpdd."
	configFileDefault = ""
	configFileKey     = keyPrefix + "configFile"
	listenHostPortKey = keyPrefix + "listenHostPort"
)

var rootCmd = &cobra.Command{
	Use:   "tftpdd",
	Short: "tftpdd is the TFTP server management daemon",
	Long: `tftpdd is the TFTP server management daemon.

Find more information at:
https://pojntfx.github.io/tftpdd/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("tftpdd")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		for {
			fmt.Println("Starting server")
			time.Sleep(time.Second * 1)
		}

		return nil
	},
}

func init() {
	var (
		configFileFlag string
		hostPortFlag   string
	)

	rootCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	rootCmd.PersistentFlags().StringVarP(&hostPortFlag, listenHostPortKey, "l", constants.TFTPDDHostPortDefault, "TCP listen host:port.")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
