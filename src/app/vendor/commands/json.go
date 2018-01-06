package commands

import (
	json "commands/json"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "json", Short: "json 格式化"}

	json.SetupPrettyCommand(cmd)

	rootCmd.AddCommand(cmd)
}
