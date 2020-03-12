package cmd

import (
	"strings"

	constants "github.com/pojntfx/tftpdd/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
)

var rootCmd = &cobra.Command{
	Use:   "tftpdctl",
	Short: "tftpdctl manages tftpdd, the TFTP server management daemon",
	Long: `tftpdctl manages tftpdd, the TFTP server management daemon.

Find more information at:
https://pojntfx.github.io/tftpdd/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("tftpd")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
