package image

import (
	"fmt"
	"github.com/spf13/cobra"
	"image/jpeg"
	"math"
	"os"
)

func SetupToJpegCommand(rootCmd *cobra.Command) {
	var quality int
	cmd := &cobra.Command{
		Use:   "to-jpg <filename>",
		Short: "轉存 jpg",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}
			img := imageRead(args[0])
			quality = int(math.Max(float64(1), math.Min(float64(100), float64(quality))))
			fn_out := fmt.Sprintf("%s-q%d.jpg", args[0], quality)
			fout, err := os.Create(fn_out)
			if err != nil {
				panic(err)
			}

			err = jpeg.Encode(fout, img, &jpeg.Options{Quality: quality})
			if err != nil {
				panic(err)
			}
			fmt.Println(fn_out)
		},
	}

	cmd.Flags().IntVarP(&quality, "quality", "q", 75, "品質, 1-100 (預設75)")
	rootCmd.AddCommand(cmd)
}
