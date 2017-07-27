package windows

import (
	common "commands/common"
	"fmt"
	ole "github.com/go-ole/go-ole"
	oleutil "github.com/go-ole/go-ole/oleutil"
	"github.com/spf13/cobra"
	"strings"
)

func SetupWord2TxtCommand(rootCmd *cobra.Command) {
	var format string
	cmd := &cobra.Command{
		Use:   "word2txt <filename>",
		Short: "透過MSWord轉存成txt",
		Run: func(cmd *cobra.Command, args []string) {
			datas := common.GetArgsOrStdIn(args)
			format = strings.ToLower(format)
			ext := ".txt"
			ext_code := 7
			if format == "pdf" {
				ext = ".pdf"
				ext_code = 17
			} else if format == "xps" {
				ext = ".xps"
				ext_code = 18
			} else if format == "html" {
				ext = ".html"
				ext_code = 8
			}
			err := ole.CoInitialize(0)
			if err != nil {
				panic(err)
			}
			defer ole.CoUninitialize()

			unknown, err := oleutil.CreateObject("Word.Application")
			if err != nil {
				panic(err)
			}
			msword, err := unknown.QueryInterface(ole.IID_IDispatch)
			if err != nil {
				panic(err)
			}
			defer oleutil.CallMethod(msword, "Quit")
			defer msword.Release()
			docs := oleutil.MustGetProperty(msword, "Documents").ToIDispatch()
			defer docs.Release()

			for _, fn := range datas {
				out_fn := fn + ext
				fmt.Println(out_fn)
				doc := oleutil.MustCallMethod(docs, "Open", fn, nil, nil, nil).ToIDispatch()
				oleutil.MustCallMethod(doc, "SaveAs2", out_fn, ext_code).ToIDispatch()
			}
		},
	}
	cmd.Flags().StringVarP(&format, "format", "f", "txt", "輸出格式(txt,pdf,xps,html")
	rootCmd.AddCommand(cmd)

}
