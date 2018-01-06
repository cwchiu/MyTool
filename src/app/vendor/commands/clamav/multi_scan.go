package clamav

import (
	"github.com/spf13/cobra"
)

func SetupMultiScanCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "multi-scan <file/directory>",
		Short: "Scan file in a standard way or scan directory (recursively) using multiple threads (to make the scanning faster on SMP machines).",
		Run:   createScanCommandHandler("MULTISCAN"),
	}
	rootCmd.AddCommand(cmd)
}
