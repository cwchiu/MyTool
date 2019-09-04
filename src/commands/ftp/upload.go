package ftp

import (
	"github.com/spf13/cobra"
	"os"
)

func SetupUploadCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "upload <local file> <remote file> ",
		Short: "upload",
		Long:  "upload test.txt ftp://test:1234@127.0.0.1/test.txt test.txt",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("required <local file> <remote file>")
			}

			info := ParseUrl(args[1])
			conn := CreateFtpClient(info)

			src, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			defer src.Close()

			err = conn.Stor(info.Path, src)
			if err != nil {
				panic(err)
			}

			// if err != nil {

			// entries, err := conn.List("/sdcard/flag.txt")
			// panic(err)
			// }

			// for _, entry := range entries {
			// fmt.Println(entry)
			// }
			// func (c *ServerConn) Retr(path string) (*Response, error)
			// func (c *ServerConn) Stor(path string, r io.Reader) error

		},
	}
	rootCmd.AddCommand(cmd)
}
