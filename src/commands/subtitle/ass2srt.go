package subtitle

import (
	"github.com/spf13/cobra"
)


func SetupAss2SrtCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "ass2srt <ass-filename>",
		Short: "ass 轉 srt",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <ass-filename>")
			}
			ass2srt(args[0])
		},
	}
	rootCmd.AddCommand(cmd)
}
