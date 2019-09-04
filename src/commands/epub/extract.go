package epub

import (
	"github.com/kapmahc/epub"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func SetupExtractCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "extract <epub> <name> <output-filename>",
		Short: "epub 取資料",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				panic("required <epub> <filename> <output-filename>")
			}
			bk, err := epub.Open(args[0])
			if err != nil {
				panic(err)
			}
			defer bk.Close()

			reader, err := bk.Open(args[1])
			if err != nil {
				panic(err)
			}
			defer reader.Close()
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(args[2], data, 0777)
			if err != nil {
				panic(err)
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
