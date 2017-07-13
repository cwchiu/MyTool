package commands

import (
	"github.com/spf13/cobra"
    barcode "commands/barcode"
)

func SetupBarCodeCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "barcode"}

	barcode.SetupQRCommand(rootCmd)
	// date.SetupT2SCommand(rootCmd)
	// date.SetupS2TCommand(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
