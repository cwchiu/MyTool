package fs

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
func SetupNlCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "nl <filename>",
		Short: "line count of file",
		Long:  `line count of file`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}

			fin, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			defer fin.Close()

			count, err := lineCounter(fin)
			if err != nil {
				panic(err)
			}

			fmt.Println(count)

		},
	}
	rootCmd.AddCommand(cmd)
}
