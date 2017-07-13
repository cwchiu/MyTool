package commands

import (
	"github.com/spf13/cobra"
    base64 "commands/base64"
)

func SetupBase64Command(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "base64"}

	base64.SetupEncodeCommand(rootCmd)
	base64.SetupDecodeCommand(rootCmd)
    
    parentCmd.AddCommand(rootCmd)
}
