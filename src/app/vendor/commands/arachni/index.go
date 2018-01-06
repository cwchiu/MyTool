package arachni

import (
	"github.com/spf13/cobra"
)

func SetupArachniCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "arachni", Short: "arachni api"}

	SetupScanStartCommand(rootCmd)
	SetupScanGetCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
