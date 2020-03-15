package cmd

import (
	"context"
	"fmt"

	constants "github.com/pojntfx/tftpdd/cmd"
	tftpdd "github.com/pojntfx/tftpdd/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"a"},
	Short:   "Apply a TFTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := tftpdd.NewTFTPDDManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.Create(ctx, &tftpdd.TFTPD{
			Device:        viper.GetString(deviceKey),
			PXEPackageURL: viper.GetString(pxepackageURLKey),
		})
		if err != nil {
			return err
		}

		fmt.Printf("TFTP server \"%s\" created\n", response.GetId())

		return nil
	},
}

func init() {
	var (
		deviceFlag        string
		pxepackageURLFlag string
	)

	applyCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.TFTPDDHostPortDefault, constants.HostPortDocs)
	applyCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	applyCmd.PersistentFlags().StringVarP(&deviceFlag, deviceKey, "d", "eth0", "Device to bind the TFTP server to")
	applyCmd.PersistentFlags().StringVarP(&pxepackageURLFlag, pxepackageURLKey, "p", "http://minio.pxepackagerd.felix.pojtinger.com/pxepackagerd/1/bin/one.pxepackage", "PXE package to use.")

	if err := viper.BindPFlags(applyCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(applyCmd)
}
