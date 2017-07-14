package commands

import (
	json "commands/json"
	"github.com/spf13/cobra"
)

func SetupJsonCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "json"}

	json.SetupPrettyCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
