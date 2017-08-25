package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	// "github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type sentencesResult struct {
	Sentences []string `json:"sentences"`
}

func SetupMoreHandlinoCommand(rootCmd *cobra.Command) {
	var lines string

	cmd := &cobra.Command{
		Use:   "more-handlino",
		Short: "中文假文產生",
		Run: func(cmd *cobra.Command, args []string) {
			request := gorequest.New()
			_, body, err := request.Get("http://more.handlino.com/sentences.json").Param("limit", "5,19").Param("n", lines).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			var result sentencesResult
			json.Unmarshal([]byte(body), &result)
			for _, line := range result.Sentences {
				fmt.Println(line)
			}
		},
	}
	cmd.Flags().StringVarP(&lines, "lines", "l", "30", "行數")
	rootCmd.AddCommand(cmd)

}
