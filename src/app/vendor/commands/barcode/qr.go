package barcode

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/spf13/cobra"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func SetupQRCommand(rootCmd *cobra.Command) {
	var name string
	var jpg bool
	cmd := &cobra.Command{
		Use:   "qr <text>",
		Short: "generate qrcode",
		Long:  `generate qrcode`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <text>")
			}
			text := args[0]
			code, err := qr.Encode(text, qr.L, qr.Auto)
			if err != nil {
				panic(err)
			}

			if text != code.Content() {
				panic("data differs")
			}

			code, err = barcode.Scale(code, 300, 300)
			if err != nil {
				panic(err)
			}

			if filepath.Ext(name) == "" {
				if jpg {
					name = name + ".jpg"
				} else {
					name = name + ".png"
				}
			}

			file, err := os.Create(name)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			if jpg {
				err = jpeg.Encode(file, code, nil)
			} else {
				err = png.Encode(file, code)
			}
			if err != nil {
				panic(err)
			}

			fmt.Println(file.Name())
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "test", "filename")
	cmd.Flags().BoolVarP(&jpg, "jpg", "j", false, "output image type")

	rootCmd.AddCommand(cmd)
}
