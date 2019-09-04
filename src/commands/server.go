package commands

import (
	server "github.com/cwchiu/MyTool/commands/server"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "server", Short: "服務器"}

	server.SetupCharGenCommand(cmd)
	server.SetupDiscardCommand(cmd)
	server.SetupEchoCommand(cmd)
	server.SetupQuoteCommand(cmd)
	server.SetupDaytimeCommand(cmd)
	server.SetupProxyCommand(cmd)
	server.SetupWebCommand(cmd)
	server.SetupFtpCommand(cmd)
	server.SetupSSHCommand(cmd)
	server.SetupTcpPortForwardCommand(cmd)
	server.SetupTunnelCommand(cmd)
	server.SetupSMTPCommand(cmd)

	rootCmd.AddCommand(cmd)
}
