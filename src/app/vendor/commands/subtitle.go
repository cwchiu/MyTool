package commands

import (
	subtitle "commands/subtitle"
	"github.com/spf13/cobra"
)

func SetupSubtitleCommand(parentCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "subtitle", Short: "字幕"}

	subtitle.SetupAss2SrtCommand(cmd)
	subtitle.SetupSrt2AssCommand(cmd)

	parentCmd.AddCommand(cmd)
}
