package commands

import (
	netcat "commands/netcat"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "netcat", Short: "netcat"}

	netcat.SetupServerCommand(cmd)
	netcat.SetupClientCommand(cmd)

	rootCmd.AddCommand(cmd)
}
