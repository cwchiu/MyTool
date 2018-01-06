package commands

import (
	date "commands/date"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "date", Short: "日期/時間"}

	date.SetupNowCommand(cmd)
	date.SetupT2SCommand(cmd)
	date.SetupS2TCommand(cmd)
	date.SetupS2LCommand(cmd)
	date.SetupL2SCommand(cmd)

	rootCmd.AddCommand(cmd)
}
