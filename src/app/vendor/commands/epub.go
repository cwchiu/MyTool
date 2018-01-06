package commands

import (
	epub "commands/epub"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "epub", Short: "epub 工具"}

	epub.SetupInfoCommand(cmd)
	epub.SetupExtractCommand(cmd)

	rootCmd.AddCommand(cmd)
}
