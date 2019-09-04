package subtitle

import (
	"github.com/spf13/cobra"
)

func SetupSrt2AssCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "srt2ass <srt-filename>",
		Short: "srt 轉 ass",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <srt-filename>")
			}
			srt2ass(args[0])
		},
	}
	rootCmd.AddCommand(cmd)
}
