package commands

import (
	"commands/url"
	"github.com/spf13/cobra"
)

func SetupUrlCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "url",
		Short: "url 相關",
	}

	url.SetupEncodeCommand(rootCmd)
	url.SetupDecodeCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
