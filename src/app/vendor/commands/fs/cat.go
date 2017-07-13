package fs

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func SetupCatCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "cat [filename]",
		Short: "show file content",
		Long:  `show file content`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}

			for _, fn := range args {
				fmt.Println(fn)
				file, err := os.Open(fn)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				buf := make([]byte, 32*1024)

				for {
					n, err := file.Read(buf)

					if n > 0 {
						fmt.Print(string(buf[:n]))
					}

					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
				}

				fmt.Println("\n")
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
