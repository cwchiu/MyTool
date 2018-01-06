package commands

import (
	barcode "commands/barcode"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "barcode", Short: "條碼"}

	barcode.SetupQRCommand(cmd)
	// date.SetupT2SCommand(rootCmd)
	// date.SetupS2TCommand(rootCmd)

	rootCmd.AddCommand(cmd)
}
