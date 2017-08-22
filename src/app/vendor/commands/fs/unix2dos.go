package fs

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

func SetupUnix2DosCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "unix2dos",
		Short: "標準輸入的資料進行 unix2dos",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				if len(line) > 0 {
					line = strings.TrimRight(line, "\r\n")
					line = strings.TrimRight(line, "\n")
					fmt.Printf("%s\r\n", line)
				}

				if err == io.EOF {
					break
				}

				if err != nil {
					panic(err)
				}
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
