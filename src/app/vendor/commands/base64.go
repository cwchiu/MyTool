package commands

import (
	base64 "commands/base64"
	"github.com/spf13/cobra"
)

func SetupBase64Command(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "base64", Short: "base64編/解碼"}

	base64.SetupEncodeCommand(rootCmd)
	base64.SetupDecodeCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
