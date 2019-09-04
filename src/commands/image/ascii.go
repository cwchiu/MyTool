package image

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
    "image"
    "bytes"
    "reflect"
    "image/color"
)

// https://gitee.com/stdupp/goasciiart/blob/master/goasciiart.go
var ASCIISTR1 = "MND8OZ$7I?+=~:,.."
var ASCIISTR2 = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "


func scaleImage(img image.Image, w int) (image.Image, int, int) {
    sz := img.Bounds()
    h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
    img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
    return img, w, h
}

func convert2Ascii(img image.Image, q, w, h int) []byte {
    var table []byte
    if q == 1 {
        table = []byte(ASCIISTR1)
    } else {
        table = []byte(ASCIISTR2)
    }
    buf := new(bytes.Buffer)
    size := uint64(len(table)-1)
    
    for i := 0; i < h; i++ {
        for j := 0; j < w; j++ {
            g := color.GrayModel.Convert(img.At(j, i))
            y := reflect.ValueOf(g).FieldByName("Y").Uint()
            pos := int(y * size / 255)
            _ = buf.WriteByte(table[pos])
        }
        _ = buf.WriteByte('\n')
    }
    return buf.Bytes()
}


func SetupToAsciiCommand(rootCmd *cobra.Command) {
	var width int
    var quantify int
	cmd := &cobra.Command{
		Use:   "ascii <filename>",
		Short: "轉存 ascii",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}
			img := imageRead(args[0])
            img, w, h := scaleImage(img, width)
            
            if quantify != 2 {
                quantify = 1
            }
            p:= convert2Ascii( img, quantify, w, h )
            fmt.Print(string(p))
		},
	}

	cmd.Flags().IntVarP(&width, "width", "w", 80, "寬度")
	cmd.Flags().IntVarP(&quantify, "quantify", "q", 1, "使用的量化表, 可用 1 或 2")
	rootCmd.AddCommand(cmd)
}
