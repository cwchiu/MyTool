package base64

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

type PreStdin struct {
	Reader  io.Reader
	checked bool
}

func (this *PreStdin) SetChecked() {
	this.checked = true
}

func (this PreStdin) Read(p []byte) (n int, err error) {
	if !this.checked {
		size := len(p)
		buf := make([]byte, 23)
		var n1 int
		n1, err = this.Reader.Read(buf)
		size2 := size - 23
		if n1 > 22 {
			s := string(buf)
			pos := strings.Index(s, ",")
			var next_pos int
			if pos > 0 {
				copy_n := copy(p, buf[pos+1:])
				size2 = size - 23 + copy_n
				next_pos = copy_n
			} else {
				copy(p, buf)
				next_pos = 23
			}
			buf2 := make([]byte, size2)
			var n2 int
			n2, err = this.Reader.Read(buf2)
			n = n2 + next_pos
			for _, b := range buf2 {
				p[next_pos] = b
				next_pos++
			}
		} else {
			copy(p, buf)
		}

		this.SetChecked()
	} else {
		n, err = this.Reader.Read(p)
	}
	return n, err
}

func SetupDecodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "decode ",
		Short: "encode stdin to base64 string",
		Long:  `encode stdin to base64 string`,
		Run: func(cmd *cobra.Command, args []string) {
			r := base64.NewDecoder(base64.StdEncoding, PreStdin{Reader: os.Stdin})
			_, err := io.Copy(os.Stdout, r)
            if err != nil {
                panic(err)
            }
		},
	}

	rootCmd.AddCommand(cmd)
}
