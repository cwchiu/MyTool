package image

import (
	"fmt"
	"github.com/spf13/cobra"
	"image/gif"
	"os"
)

func SetupToGifCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "to-gif <filename>",
		Short: "轉存 gif",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}
			img := imageRead(args[0])
			fn_out := fmt.Sprintf("%s.gif", args[0])
			fout, err := os.Create(fn_out)
			if err != nil {
				panic(err)
			}
			var opt gif.Options
			opt.NumColors = 256
			err = gif.Encode(fout, img, &opt)
			if err != nil {
				panic(err)
			}
			fmt.Println(fn_out)
		},
	}
	rootCmd.AddCommand(cmd)
}
