package web

import (
	"fmt"
	// "net/url"
    "github.com/parnurzeal/gorequest"
    "github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)


func SetupYoudaoDictCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "youdao-dict <word>",
		Short: "有道字典詞語翻譯",
		Long:  `有道字典詞語翻譯`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("required <word>")
			}
            word := args[0]
            request := gorequest.New()
			resp, _, err := request.Get("http://dict.youdao.com/search").Param("keyfrom","dict.index").Param("q", word).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
            // fmt.Println(body)
            doc, err2 := goquery.NewDocumentFromResponse(resp)
            if err2 != nil {
                panic(err2)
            }
            doc.Find("#results-contents > #phrsListTab > .trans-container li").Each(func(i int, s *goquery.Selection) {
                    fmt.Println(s.Text())
              })
		},
	}
	rootCmd.AddCommand(cmd)

}
