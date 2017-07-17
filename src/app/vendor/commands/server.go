package commands

import (
	server "commands/server"
	"github.com/spf13/cobra"
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
	server.SetupFtpCommand(rootCmd)
	server.SetupSSHCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
