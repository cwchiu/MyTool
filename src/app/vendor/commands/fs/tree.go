package fs

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	// "path/filepath"
)

func s_printf(filename string, depth int) {
	for depth > 0 {
		depth--
		fmt.Printf("|  ")
	}
	fmt.Printf("|--")
	fmt.Println(filename)
}

func s_dirwalk(dirname string, depth int, method func(v string, d int)) {
	f, _ := os.Open(dirname)
	defer f.Close()
	fis, _ := f.Readdirnames(0)
	for _, fn := range fis {
		if fn == "." || fn == ".." {
			continue
		}

		s_printf(fn, depth)

		method(dirname+"/"+fn, depth+1)

		//fmt.Println(fn)
	}
}

func listdirtree(dirname string, depth int) {
	f, err := os.Open(dirname)
	if err != nil {
		fmt.Println("can't read file " + dirname)
		return
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("can't read file " + dirname)
		return
	}

	if fi.IsDir() {
		//fmt.Println("Dir")
		s_dirwalk(dirname, depth, listdirtree)
	} else {
		//fmt.Println("File")
	}
}

// http://www.oschina.net/code/snippet_614309_14170
func SetupTreeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "tree [path]",
		Short: "list files",
		Long:  `list files`,
		Run: func(cmd *cobra.Command, args []string) {
			src := "."
			if len(args) > 0 {
				src = args[0]
			}
			listdirtree(src, 0)
		},
	}
	rootCmd.AddCommand(cmd)
}
