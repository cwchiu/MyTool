package clamav

import (
	"github.com/spf13/cobra"
)

func SetupAllMatchScanCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "all-match-scan <file>",
		Short: "works just like SCAN except that it sets a mode where scanning continues after finding a match within a file.",
		Run:   createScanCommandHandler("ALLMATCHSCAN"),
	}
	rootCmd.AddCommand(cmd)
}
