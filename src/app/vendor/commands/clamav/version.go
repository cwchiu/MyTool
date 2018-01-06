package clamav

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetupVersionCommand(rootCmd *cobra.Command) {
	var server string

	cmd := &cobra.Command{
		Use:   "version",
		Short: "ClamAV Version",
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				panic("required -s <ClamAV server host> ")
			}
			inst, err := CreateClamAV(server)
			if err != nil {
				panic(err)
			}
			ver, err := inst.Version()
			if err != nil {
				panic(err)
			}
			fmt.Println(ver)

		},
	}
	cmd.Flags().StringVarP(&server, "server", "s", "192.168.99.100:3310", "ClamAV server host")
	rootCmd.AddCommand(cmd)
}
