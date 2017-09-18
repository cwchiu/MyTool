package commands

import (
	date "commands/date"
	"github.com/spf13/cobra"
)

func SetupDateCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "date", Short: "日期/時間"}

	date.SetupNowCommand(rootCmd)
	date.SetupT2SCommand(rootCmd)
	date.SetupS2TCommand(rootCmd)
	date.SetupS2LCommand(rootCmd)
	date.SetupL2SCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
