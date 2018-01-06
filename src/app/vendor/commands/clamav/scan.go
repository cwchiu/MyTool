package clamav

import (
	"github.com/spf13/cobra"
)

func SetupScanCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "scan <file/directory>",
		Short: "Scan a file or directory given by filename and stop on first virus or error found.",
		Run:   createScanCommandHandler("SCAN"),
	}
	rootCmd.AddCommand(cmd)
}
