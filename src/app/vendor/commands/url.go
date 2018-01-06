package commands

import (
	"commands/url"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "url",
		Short: "url 相關",
	}

	url.SetupEncodeCommand(cmd)
	url.SetupDecodeCommand(cmd)

	rootCmd.AddCommand(cmd)
}
