package commands

import (
	"github.com/spf13/cobra"
    server "commands/server"
)


func SetupServerCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "server"}

	server.SetupCharGenCommand(rootCmd)
	server.SetupDiscardCommand(rootCmd)
	server.SetupEchoCommand(rootCmd)
	server.SetupQuoteCommand(rootCmd)
	server.SetupDaytimeCommand(rootCmd)
	server.SetupProxyCommand(rootCmd)
	server.SetupStaticCommand(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
