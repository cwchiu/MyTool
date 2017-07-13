package fs

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
)

func SetupRmCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "rm [filename]",
		Short: "delete file",
		Long:  `delete file`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}

			for _, fn := range args {
				err := os.Remove(fn)
				if err != nil {
					panic(err)
				}
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
