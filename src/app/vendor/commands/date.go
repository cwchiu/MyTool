package commands

import (
	date "commands/date"
	"github.com/spf13/cobra"
)

func SetupDateCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "date"}

	date.SetupNowCommand(rootCmd)
	date.SetupT2SCommand(rootCmd)
	date.SetupS2TCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
