package commands

import (
	barcode "commands/barcode"
	"github.com/spf13/cobra"
)

func SetupBarCodeCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "barcode"}

	barcode.SetupQRCommand(rootCmd)
	// date.SetupT2SCommand(rootCmd)
	// date.SetupS2TCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
