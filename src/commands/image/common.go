package image

import (
	"bufio"
	"golang.org/x/image/webp"
	_image "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

func imageRead(fn string) _image.Image {
	fin, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	r := bufio.NewReader(fin)
	buf := make([]byte, 512)
	n, err := r.Read(buf)
	if err != nil {
		panic(err)
	}
	// fmt.Println(n)
	content_type := http.DetectContentType(buf[:n])
	_, err = fin.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	var img _image.Image
	switch content_type {
	case "image/webp":
		img, err = webp.Decode(fin)
		if err != nil {
			panic(err)
		}
	case "image/jpeg":
		img, err = jpeg.Decode(fin)
		if err != nil {
			panic(err)
		}
	case "image/png":
		img, err = png.Decode(fin)
		if err != nil {
			panic(err)
		}
	case "image/gif":
		img, err = gif.Decode(fin)
		if err != nil {
			panic(err)
		}
	default:
		panic("not support " + content_type)
	}

	return img
}
