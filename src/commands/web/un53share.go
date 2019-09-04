package web

import (
	"fmt"
    "net/http"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
    "github.com/PuerkitoBio/goquery"
    "regexp"
)

func SetupUn53shareCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "un53share <url>",
		Short: "un53share",
		Long:  `un53share`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("need an url")
			}
            url := args[0] // "http://53share.com/lA3Ql"
            matched, regex_err := regexp.MatchString(`(?i)^http://53share.com/\w+`, url)
            if matched == false || regex_err != nil {
                fmt.Printf("%s not support\n", url)
                return
            }
			request := gorequest.New()
            
			resp, _, err := request.Get(url).RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
                        return http.ErrUseLastResponse
                    }).Set("Referer", url).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
            doc, err2 := goquery.NewDocumentFromResponse(resp); 
            if err2 != nil {
                panic(err2)
            }
            
            href, exists := doc.Find("a.redirect[rel=nofollow]").Attr("href")
            if !exists {
                fmt.Println("not found")
                return
            // }else{
                // fmt.Println(href)
            }
            
            resp, _, err = request.Head(href).RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
                        return http.ErrUseLastResponse
                    }).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").End()
			if err != nil {
				panic(err)
			}
            fmt.Println(resp.Header["Location"][0])
		},
	}

	rootCmd.AddCommand(cmd)

}
