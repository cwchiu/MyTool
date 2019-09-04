package commands

import (
	base64 "github.com/cwchiu/MyTool/commands/base64"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "base64", Short: "base64編/解碼"}

	base64.SetupEncodeCommand(cmd)
	base64.SetupDecodeCommand(cmd)

	rootCmd.AddCommand(cmd)
}
