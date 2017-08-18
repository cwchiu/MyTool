package commands

import (
	pid "commands/pid"
	"github.com/spf13/cobra"
)

func SetupPidCommand(parentCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "pid", Short: "身分證產生"}

	pid.SetupTaiwanPidCommand(cmd)
	pid.SetupChinaPidCommand(cmd)

	parentCmd.AddCommand(cmd)
}
