package fs

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
)

func SetupMvCommand(rootCmd *cobra.Command) {
	var src string
	var dst string
	cmd := &cobra.Command{
		Use:   "mv",
		Short: "move file from srouce to target",
		Long:  `move file from srouce to target`,
		Run: func(cmd *cobra.Command, args []string) {
			err := os.Rename(src, dst)
			if err != nil {
				panic(err)
			}

		},
	}

	cmd.Flags().StringVarP(&src, "source", "s", "", "source filename")
	cmd.Flags().StringVarP(&dst, "target", "t", "", "target filename")
	rootCmd.AddCommand(cmd)
}
