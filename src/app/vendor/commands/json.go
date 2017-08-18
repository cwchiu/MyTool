package commands

import (
	json "commands/json"
	"github.com/spf13/cobra"
)

func SetupJsonCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "json", Short: "json 格式化"}

	json.SetupPrettyCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
