package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	// "net/url"
	// "sort"
	// "strings"
)

type apiresult struct {
	Text             []string            `json:"text"`
	ErrorCode        int                 `json:"code"`
	Message          string              `json:"message"`
}

func SetupYandexTranslateCommand(rootCmd *cobra.Command) {
	var key string
	var lang string

	cmd := &cobra.Command{
		Use:   "yandex-trans <word>",
		Short: "Yandex 翻譯服務",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <word>")
			}
            
            if len(key) == 0 {
                panic("required -k, https://tech.yandex.com/keys/get/?service=trnsl")
            }
            
			word := args[0]
            request := gorequest.New()
			_, body, err := request.Get("https://translate.yandex.net/api/v1.5/tr.json/translate").Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).Param("text", word).Param("key", key).Param("lang", lang).Param("format", "plain").End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			var result apiresult
			json.Unmarshal([]byte(body), &result)
			// fmt.Println(result)
			if result.ErrorCode != 200 {
				fmt.Println(result.Message)
				return
			}
			fmt.Println(result.Text[0])
		},
	}

	cmd.Flags().StringVarP(&key, "key", "k", "", "API key, https://tech.yandex.com/keys/get/?service=trnsl")
	cmd.Flags().StringVarP(&lang, "lang", "l", "ja", "轉換後的語言, https://tech.yandex.com/translate/doc/dg/concepts/api-overview-docpage/")
	rootCmd.AddCommand(cmd)

}
