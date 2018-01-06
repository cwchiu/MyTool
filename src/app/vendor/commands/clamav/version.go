package clamav

import (
	"github.com/spf13/cobra"
)

func SetupVersionCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "ClamAV Version",
		Run:   createNoArgsCommandHandler("VERSION"),
	}
	rootCmd.AddCommand(cmd)
}
