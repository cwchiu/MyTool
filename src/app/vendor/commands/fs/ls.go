package fs

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func SetupLsCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "ls [path]",
		Short: "list files",
		Long:  `list files`,
		Run: func(cmd *cobra.Command, args []string) {
			src := "."
			if len(args) > 0 {
				src = args[0]
			}
			filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					fmt.Printf("[%s]\n", path)
				} else {
					fmt.Println(path)
				}
				return nil
			})
		},
	}
	rootCmd.AddCommand(cmd)
}
