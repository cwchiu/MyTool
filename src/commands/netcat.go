package commands

import (
	netcat "github.com/cwchiu/MyTool/commands/netcat"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "netcat", Short: "netcat"}

	netcat.SetupServerCommand(cmd)
	netcat.SetupClientCommand(cmd)

	rootCmd.AddCommand(cmd)
}
