package barcode

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

func SetupEncodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "encode ",
		Short: "encode stdin to base64 string",
		Long:  `encode stdin to base64 string`,
		Run: func(cmd *cobra.Command, args []string) {
			r := bufio.NewReader(os.Stdin)
			buf := make([]byte, 0, 4*1024)
			content_type := ""
			var b bytes.Buffer
			w := base64.NewEncoder(base64.StdEncoding, &b)
			for {
				n, err := r.Read(buf[:cap(buf)])
				buf = buf[:n]
				if n == 0 {
					if err == nil {
						continue
					}
					if err == io.EOF {
						break
					}
					panic(err)
				}
				if err != nil && err != io.EOF {
					panic(err)
				}
				if content_type == "" {
					content_type = http.DetectContentType(buf)
					// fmt.Println(content_type)
				}
				w.Write(buf)

			}
			w.Close()
			if content_type == "image/png" {
				fmt.Print("data:image/png;base64,")
			} else if content_type == "image/jpeg" {
				fmt.Print("data:image/jpeg;base64,")
			}
			fmt.Println(b.String())
		},
	}

	rootCmd.AddCommand(cmd)
}
