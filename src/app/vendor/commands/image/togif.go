package image

import (
	"fmt"
	"github.com/ritchie46/GOPHY/img2gif"
	"github.com/spf13/cobra"
	"image"
	"image/gif"
	"os"
)

func SetupToGifCommand(rootCmd *cobra.Command) {
	var merge bool
	var fps int
	cmd := &cobra.Command{
		Use:   "to-gif <filename>",
		Short: "轉存 gif",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}

			im := []image.Image{}
			for _, fn := range args {
				im = append(im, imageRead(fn))
			}

			if merge {
				im_p := img2gif.EncodeImgPaletted(&im)
				err := img2gif.WriteGif(&im_p, 100/fps, args[0]+"-m.gif")
				if err != nil {
					panic(err)
				}
			} else {
				for i, fn := range args {
					fn_out := fmt.Sprintf("%s.gif", fn)
					fout, err := os.Create(fn_out)
					if err != nil {
						panic(err)
					}
					var opt gif.Options
					opt.NumColors = 256
					err = gif.Encode(fout, im[i], &opt)
					if err != nil {
						panic(err)
					}
					fmt.Println(fn_out)
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&merge, "merge", "m", false, "合併成動畫")
	cmd.Flags().IntVarP(&fps, "fps", "f", 1, "fps")
	rootCmd.AddCommand(cmd)
}
