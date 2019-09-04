package mp3

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetupPlayCommand(rootCmd *cobra.Command) {
	var repeat int
	var shuffle bool
	cmd := &cobra.Command{
		Use:   "play <filename>",
		Short: "播放mp3",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not support")
		},
	}

	cmd.Flags().IntVarP(&repeat, "repeat", "r", 1, "重複次數, 小於 1 表示無限重複")
	cmd.Flags().BoolVarP(&shuffle, "shuffle", "s", false, "亂數播放")
	rootCmd.AddCommand(cmd)
}
