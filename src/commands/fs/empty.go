package fs

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
)

func SetupEmptyCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "empty [filename]",
		Short: "create empty file",
		Long:  `create empty file`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}

			for _, fn := range args {
				newFile, err := os.Create(fn)
				if err != nil {
					panic(err)
				}
				defer newFile.Close()
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
