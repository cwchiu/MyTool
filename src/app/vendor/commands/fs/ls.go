package fs

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func SetupLsCommand(rootCmd *cobra.Command) {
	var recursive bool
	cmd := &cobra.Command{
		Use:   "ls <path>",
		Short: "檔案清單",
		Run: func(cmd *cobra.Command, args []string) {
			src := "."
			if len(args) > 0 {
				src = args[0]
			}

			if recursive == false {
				matches, err := filepath.Glob(src)
				if err != nil {
					panic(err)
				}
				for _, fn := range matches {
					fmt.Println(fn)
				}
			} else {
				ext := filepath.Ext(src)
				if ext != "" {
					if filepath.Base(src) != "*"+ext {
						if _, err := os.Stat(src); err == nil {
							fmt.Println(src)
						}
						return
					}

					src = filepath.Dir(src)
				}

				filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						// fmt.Println(err)
						return nil
					}
					switch {
					case ext != "":
						// fmt.Println(path)
						if filepath.Ext(path) == ext {
							fmt.Println(path)
						}
					case info.IsDir():
						fmt.Printf("[%s]\n", path)
					default:
						fmt.Println(path)
					}
					return nil
				})
			}
		},
	}
	cmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "遞迴查找子目錄")
	rootCmd.AddCommand(cmd)
}
