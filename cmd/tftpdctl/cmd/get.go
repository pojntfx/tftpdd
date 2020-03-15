package cmd

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/gosuri/uitable"
	constants "github.com/pojntfx/tftpdd/cmd"
	TFTPDD "github.com/pojntfx/tftpdd/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var getCmd = &cobra.Command{
	Use:     "get [id]",
	Aliases: []string{"g"},
	Short:   "Get one or all TFTP server(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := TFTPDD.NewTFTPDDManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if len(args) < 1 {
			response, err := client.List(ctx, &TFTPDD.TFTPDManagerListArgs{})
			if err != nil {
				return err
			}

			table := uitable.New()
			table.AddRow(
				"ID",
				"LISTEN ADDRESS",
				"STATUS")

			for _, DHCPD := range response.GetTFTPDs() {
				table.AddRow(
					DHCPD.GetId(),
					DHCPD.GetListenAddress(),
					DHCPD.GetStatus())
			}

			fmt.Println(table)

			return nil
		}

		response, err := client.Get(ctx, &TFTPDD.TFTPDId{
			Id: args[0],
		})
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(&response)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

		return nil
	},
}

func init() {
	getCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.TFTPDDHostPortDefault, constants.HostPortDocs)

	if err := viper.BindPFlags(getCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(getCmd)
}
