package commands

import (
	"github.com/spf13/cobra"
    date "commands/date"
)


func SetupDateCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "date"}

	date.SetupNowCommand(rootCmd)
	date.SetupT2SCommand(rootCmd)
	date.SetupS2TCommand(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
