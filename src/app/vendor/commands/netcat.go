package commands

import (
	netcat "commands/netcat"
	"github.com/spf13/cobra"
)

func SetupNetcatCommand(parentCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "netcat", Short: "netcat"}

	netcat.SetupServerCommand(cmd)
	netcat.SetupClientCommand(cmd)

	parentCmd.AddCommand(cmd)
}
