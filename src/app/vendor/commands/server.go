package commands

import (
	server "commands/server"
	"github.com/spf13/cobra"
)

func SetupServerCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "server", Short: "服務器"}

	server.SetupCharGenCommand(rootCmd)
	server.SetupDiscardCommand(rootCmd)
	server.SetupEchoCommand(rootCmd)
	server.SetupQuoteCommand(rootCmd)
	server.SetupDaytimeCommand(rootCmd)
	server.SetupProxyCommand(rootCmd)
	server.SetupWebCommand(rootCmd)
	server.SetupFtpCommand(rootCmd)
	server.SetupSSHCommand(rootCmd)
	server.SetupTcpPortForwardCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
