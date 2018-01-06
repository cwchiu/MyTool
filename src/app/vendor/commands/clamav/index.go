package clamav

import (
	"github.com/spf13/cobra"
)

func SetupCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "ClamAV", Short: "ClamAV api"}

	SetupVersionCommand(rootCmd)
	SetupScanCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
