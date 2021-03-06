package commands

import (
	"github.com/spf13/cobra"
    "commands/proxypool"
)


func init() {
	cmd := &cobra.Command{
		Use:   "proxy-pool",
		Short: "ProxyPool Service",
	}

	// setupCheckCommand(rootCmd)
	proxypool.SetupRunCommand(cmd)

	rootCmd.AddCommand(cmd)
}
