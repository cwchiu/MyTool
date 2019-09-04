package commands

import (
	mp3 "github.com/cwchiu/MyTool/commands/mp3"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "mp3", Short: "mp3 播放/資訊"}

	mp3.SetupPlayCommand(cmd)
	mp3.SetupInfoCommand(cmd)

	rootCmd.AddCommand(cmd)
}
