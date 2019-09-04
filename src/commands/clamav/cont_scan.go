package clamav

import (
	"github.com/spf13/cobra"
)

func SetupContScanCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "cont-scan <file/directory>",
		Short: "Scan file or directory (recursively) with archive support enabled and don't stop the scanning when a virus is found.",
		Run:   createScanCommandHandler("CONTSCAN"),
	}
	rootCmd.AddCommand(cmd)
}
