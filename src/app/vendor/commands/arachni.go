package commands

import (
    "commands/arachni"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "arachni", Short: "arachni api"}

	arachni.SetupScanStartCommand(cmd)
	arachni.SetupScanGetCommand(cmd)

	rootCmd.AddCommand(cmd)
}
