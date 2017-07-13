package fs

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
)

func SetupTruncatCommand(rootCmd *cobra.Command) {
	var size int64
	cmd := &cobra.Command{
		Use:   "truncat [filename]",
		Short: "show file info",
		Long:  `show file info`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}

			for _, fn := range args {
				err := os.Truncate(fn, size)
				if err != nil {
					panic(err)
				}
			}

		},
	}

	cmd.Flags().Int64VarP(&size, "size", "s", 1024, "truncate size")
	rootCmd.AddCommand(cmd)
}
