package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type GAnswer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int
	Data string `json:"data"`
}

type GQuestion struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type GDNSResults struct {
	Status   int
	TC       bool
	RD       bool
	AD       bool
	CD       bool
	Comment  string
	Question GQuestion
	Answer   []GAnswer
}

func SetupGoogleDnsResolveCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "dns-resolve <uri>",
		Short: "DNS名稱解析",
		Long:  `DNS名稱解析`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <uri>")
			}
			uri := args[0]
			request := gorequest.New()
			_, body, err := request.Get("https://dns.google.com/resolve").Param("name", uri).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			var result GDNSResults
			json.Unmarshal([]byte(body), &result)
			// fmt.Println(result)
			if result.Status != 0 {
				fmt.Println("fail")
				return
			}
			if len(result.Answer) < 1 {
				fmt.Println("not found")
				return
			}
			for _, data := range result.Answer {
				fmt.Printf("Name: %v\n", data.Name)
				fmt.Printf("IP: %v\n", data.Data)
				fmt.Printf("TTL: %v\n", data.TTL)
				fmt.Printf("Type: %v\n", data.Type)
				fmt.Println("-------------")
			}

		},
	}
	rootCmd.AddCommand(cmd)

}
