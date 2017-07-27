package image

import (
	"fmt"
	"github.com/spf13/cobra"
	"image/png"
	"os"
)

func SetupToPngCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "to-png <filename>",
		Short: "轉存 png",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}
			img := imageRead(args[0])
			fn_out := fmt.Sprintf("%s.png", args[0])
			fout, err := os.Create(fn_out)
			if err != nil {
				panic(err)
			}

			err = png.Encode(fout, img)
			if err != nil {
				panic(err)
			}
			fmt.Println(fn_out)
		},
	}
	rootCmd.AddCommand(cmd)
}
