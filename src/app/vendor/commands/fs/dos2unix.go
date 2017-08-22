package fs

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

func SetupDos2UnixCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "dos2unix",
		Short: "標準輸入的資料進行 dos2unix",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				if len(line) > 0 {
					// fmt.Println(line)
					// datas = append(datas, strings.TrimSpace(line))
					fmt.Print(strings.Replace(line, "\r\n", "\n", -1))
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
