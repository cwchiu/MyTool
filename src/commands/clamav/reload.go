package clamav

import (
	"github.com/spf13/cobra"
)

func SetupReloadCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "reload",
		Short: "Force Clamd to reload signature database",
		Run:   createNoArgsCommandHandler("RELOAD"),
	}
	rootCmd.AddCommand(cmd)
}
