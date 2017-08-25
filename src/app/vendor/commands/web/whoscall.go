package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"strings"
)

func SetupWhosCallCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "whoscall <number>",
		Short: "查號碼來源",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("required <number>")
			}

			request := gorequest.New()
			resp, _, err := request.Get(fmt.Sprintf("https://whoscall.com/zh-TW/tw/%s/", args[0])).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			doc, err2 := goquery.NewDocumentFromResponse(resp)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println(strings.TrimSpace(doc.Find(".number-info .number-info__name").First().Text()))
			fmt.Println(strings.TrimSpace(doc.Find(".number-info .number-info__subinfo").First().Text()))
		},
	}
	rootCmd.AddCommand(cmd)

}
