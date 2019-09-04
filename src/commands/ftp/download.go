package ftp

import (
	"github.com/spf13/cobra"
	// "fmt"
	"io"
	"os"
)

func SetupDownloadCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "download <remote file> <local file>",
		Short: "download",
		Long:  "download ftp://test:1234@127.0.0.1/test.txt test.txt",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("required <remote file> <local file>")
			}
			info := ParseUrl(args[0])
			conn := CreateFtpClient(info)
			resp, err := conn.Retr(info.Path)
			if err != nil {
				panic(err)
			}

			target, err := os.Create(args[1])
			if err != nil {
				panic(err)
			}
			defer target.Close()

			_, err = io.Copy(target, resp)
			if err != nil {
				panic(err)
			}
		},
	}
	rootCmd.AddCommand(cmd)
}
