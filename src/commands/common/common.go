package common

import (
    "bufio"
    "os"
    "strings"
    "io"
)

func GetArgsOrStdIn(args []string) []string {
	var datas []string
	if len(args) < 1 {
		reader := bufio.NewReader(os.Stdin)

		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				// fmt.Println(line)
				datas = append(datas, strings.TrimSpace(line))
			}

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}
		}
	} else {
		datas = args
	}
    
    return datas
}
