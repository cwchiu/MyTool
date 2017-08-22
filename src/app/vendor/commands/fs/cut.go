package fs

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strconv"
	"strings"
)

func SetupCutCommand(rootCmd *cobra.Command) {
	var delimiter string
	var fields string

	cmd := &cobra.Command{
		Use:   "cut",
		Short: "剪裁資料",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			var parts []string
			fields_value, err := strconv.Atoi(fields)
			if err != nil {
				panic(err)
			}
			fields_value -= 1

			if fields_value < 0 {
				panic("fields 必須大於0")
			}

			for {
				line, err := reader.ReadString('\n')
				if len(line) > 0 {
					line = strings.TrimRight(line, "\r\n")
					line = strings.TrimRight(line, "\n")
					if len(delimiter) > 0 {
						parts = strings.Split(line, delimiter)
					} else {
						parts = []string{line}
					}
					// fmt.Printf(">%v<\n", parts)
					if fields_value < len(parts) {
						fmt.Println(parts[fields_value])
					}
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
	cmd.Flags().StringVarP(&fields, "fields", "f", "1", "欄位")
	cmd.Flags().StringVarP(&delimiter, "delimiter", "d", "", "分隔符號")
	rootCmd.AddCommand(cmd)
}
