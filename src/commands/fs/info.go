package fs

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func SetupInfoCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "info [filename]",
		Short: "show file info",
		Long:  `show file info`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}

			for _, fn := range args {
				fileInfo, err := os.Stat(fn)
				if err != nil {
					panic(err)
				}

				fmt.Println("File name:", fileInfo.Name())
				fmt.Println("Size in bytes:", fileInfo.Size())
				fmt.Println("Permissions:", fileInfo.Mode())
				fmt.Println("Last modified:", fileInfo.ModTime())
				fmt.Println("Is Directory: ", fileInfo.IsDir())
				fmt.Printf("System interface type: %T\n", fileInfo.Sys())
				fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
