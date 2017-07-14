package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"net/url"
	"strings"
)

func SetupGenChineseNameCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "gen-cht-name",
		Short: "隨機產生中文姓名",
		Run: func(cmd *cobra.Command, args []string) {
			request := gorequest.New()

			form := url.Values{}
			// form.Add("break", "4")
			// form.Add("name_count", "100")
			resp, _, err := request.Post("http://www.richyli.com/name/index.asp").Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").
				Type("urlencoded").Send(form.Encode()).SetDebug(true).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			doc, err2 := goquery.NewDocumentFromResponse(resp)
			if err2 != nil {
				panic(err2)
			}
			ret, _ := doc.Find("table tr:nth-of-type(3) > td:nth-of-type(1)").First().Html()
			str2, _, _ := transform.String(traditionalchinese.Big5.NewDecoder(), ret)
			for _, name := range strings.Split(str2, "、") {
				name = strings.TrimSpace(name)
				if strings.Index(name, "名單結束") != -1 {
					break
				}

				if name == "" {
					continue
				}

				fmt.Println(name)
			}
		},
	}
	rootCmd.AddCommand(cmd)

}
