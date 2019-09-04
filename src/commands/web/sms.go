package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
    "strings"
)

func SetupSMSCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "sms",
		Short: "可用的SMS列表",
		Run: func(cmd *cobra.Command, args []string) {
			request := gorequest.New()

			resp, _, err := request.Get("http://getfreesmsnumber.com/").Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			doc, err2 := goquery.NewDocumentFromResponse(resp)
			if err2 != nil {
				panic(err2)
			}
			doc.Find(".page-header > div > div.nostyle").Each(func(i int, s *goquery.Selection) {
                lines := strings.Split(s.Text(), "\n")
                fmt.Println( "# " + strings.TrimSpace(lines[2]) )
                fmt.Println( "* " + strings.TrimSpace(lines[1]))
                fmt.Println( "* http://getfreesmsnumber.com" + s.Find("a").AttrOr("href", "") )
                fmt.Println()
            })
		},
	}
	rootCmd.AddCommand(cmd)

}
