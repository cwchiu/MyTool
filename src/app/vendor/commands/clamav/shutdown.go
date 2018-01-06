package clamav

import (
	"github.com/spf13/cobra"
)

func SetupShutdownCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "shutdown",
		Short: "Force Clamd to shutdown and exit",
		Run:   createNoArgsCommandHandler("SHUTDOWN"),
	}
	rootCmd.AddCommand(cmd)
}
