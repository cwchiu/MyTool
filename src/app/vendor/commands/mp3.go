package commands

import (
	mp3 "commands/mp3"
	"github.com/spf13/cobra"
)

func SetupMp3Command(parentCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "mp3", Short: "mp3 播放/資訊"}

	mp3.SetupPlayCommand(cmd)
	mp3.SetupInfoCommand(cmd)

	parentCmd.AddCommand(cmd)
}
