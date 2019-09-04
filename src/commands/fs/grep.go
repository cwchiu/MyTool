package fs

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"regexp"
	// "strconv"
	"strings"
)

func SetupGrepCommand(rootCmd *cobra.Command) {
	var enable_regex bool
	var ignore_case bool

	cmd := &cobra.Command{
		Use:   "grep <string>",
		Short: "過濾資料",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("required <string>")
			}
			keyword := args[0]
			if ignore_case {
				keyword = strings.ToLower(keyword)
			}
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				if len(line) > 0 {
					// line = strings.TrimRight(line, "\r\n")
					// line = strings.TrimRight(line, "\n")
					// if len(delimiter) > 0 {
					// parts = strings.Split(line, delimiter)
					// } else {
					// parts = []string{line}
					// }
					// fmt.Printf(">%v<\n", parts)
					// if fields_value < len(parts) {
					// fmt.Println(parts[fields_value])
					// }
					text := line
					if ignore_case {
						text = strings.ToLower(line)
					}

					if enable_regex {
						matched, err := regexp.Match(".*"+keyword+".*", []byte(text))
						if err == nil && matched {
							fmt.Print(line)
						}
					} else {
						if strings.Contains(text, keyword) {
							fmt.Print(line)
						}
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
	cmd.Flags().BoolVarP(&enable_regex, "regex", "r", false, "是否啟用Regex")
	cmd.Flags().BoolVarP(&ignore_case, "ignore-case", "i", false, "忽略大小寫")
	rootCmd.AddCommand(cmd)
}
