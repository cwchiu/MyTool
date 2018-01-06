package clamav

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetupScanCommand(rootCmd *cobra.Command) {
	var server string

	cmd := &cobra.Command{
		Use:   "scan <file>",
		Short: "啟動掃描",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <file>")
			}

			if server == "" {
				panic("required -s <ClamAV server host> ")
			}
			inst, err := CreateClamAV(server)
			if err != nil {
				panic(err)
			}
			scan_result, err := inst.Scan(args[0])
			if err != nil {
				panic(err)
			}
			fmt.Println(scan_result)
		},
	}
	cmd.Flags().StringVarP(&server, "server", "s", "192.168.99.100:3310", "ClamAV server host")
	rootCmd.AddCommand(cmd)
}
