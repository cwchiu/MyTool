package clamav

import (
	"github.com/spf13/cobra"
)

func SetupPingCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "ping",
		Short: "Send a PING to the clamav server, which should reply by a PONG.",
		Run:   createNoArgsCommandHandler("PING"),
	}
	rootCmd.AddCommand(cmd)
}
