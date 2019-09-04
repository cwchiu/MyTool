package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

func SetupBabelGenCommand(rootCmd *cobra.Command) {
	var keyword string
	cmd := &cobra.Command{
		Use:   "babel-generator",
		Short: "英文假文產生器",
		Run: func(cmd *cobra.Command, args []string) {
			request := gorequest.New()
			resp, _, err := request.Get("http://babel-generator.herokuapp.com/").Param("keyword", keyword).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			doc, err2 := goquery.NewDocumentFromResponse(resp)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println(doc.Find(".essay-contents").First().Text())
		},
	}
	cmd.Flags().StringVarP(&keyword, "keyword", "k", "health", "關鍵字")
	rootCmd.AddCommand(cmd)

}
