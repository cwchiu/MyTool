package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
    "regexp"
	"strings"
)

func SetupZipCodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "zipcode <address>",
		Short: "查地址對應郵遞區號",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("required <address>")
			}

			request := gorequest.New()
			resp, _, err := request.Get("http://zipko.info/?f=Query").Param("addr", args[0]).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			doc, err2 := goquery.NewDocumentFromResponse(resp)
			if err2 != nil {
				panic(err2)
			}
			result := strings.TrimSpace(doc.Find(".fulladdrSpan").First().Text())
            if len(result) > 0 {
                fmt.Println(result)
                return
            }
            fmt.Println("可能的地址有")
            result = strings.TrimSpace(doc.Find(".resultDiv").First().Text())
            // fmt.Println(result)
            re := regexp.MustCompile("完整地址：(.*)")
            i := 0
            for _, addr := range(re.FindAllStringSubmatch(result, -1)) {
                fmt.Println(addr[1])
                i += 1
            }
            
            if(i==0){
                fmt.Println("無")
            }
			// fmt.Println(strings.TrimSpace(doc.Find(".number-info .number-info__subinfo").First().Text()))
		},
	}
	rootCmd.AddCommand(cmd)

}
