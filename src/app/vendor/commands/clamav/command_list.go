package clamav

import (
	"github.com/spf13/cobra"
)

func SetupCommandListCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "command-list",
		Short: "List Support command",
		Run:   createNoArgsCommandHandler("nVERSIONCOMMANDS"),
	}
	rootCmd.AddCommand(cmd)
}
