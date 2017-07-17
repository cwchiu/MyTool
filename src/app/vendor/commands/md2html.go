package commands

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	// "time"
	"github.com/skratchdot/open-golang/open"
	"path/filepath"
	// "github.com/russross/blackfriday"
	"github.com/shurcooL/github_flavored_markdown"
)

func tempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

func SetupMD2HtmlCommand(rootCmd *cobra.Command) {
	var flagvar bool
	cmd := &cobra.Command{
		Use:   "md2html",
		Short: "markdown to html",
		Long:  `markdown to html`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}

			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				return
			}

			source, _ := ioutil.ReadFile(args[0])
			// output := blackfriday.MarkdownBasic(source)
			// fmt.Println( string(output) )
			// file, _ := ioutil.TempFile(os.TempDir(), "md2htm_")
			file := tempFileName("md2html", ".html")
			// defer os.Remove(file)
			output := github_flavored_markdown.Markdown(source)
			fout, err := os.OpenFile(file, os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			defer fout.Close()
			fout.WriteString(`<html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head><body>`)
			fout.WriteString(string(output))
			fout.WriteString("</body></html>")
			if flagvar {
				open.Start(file)
			}
			fmt.Println(file)
			// time.Sleep(10)
		},
	}
	cmd.Flags().BoolVarP(&flagvar, "open", "o", false, "open file")
	rootCmd.AddCommand(cmd)
}
