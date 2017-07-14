package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"net/url"
	"sort"
	"strings"
)

type TranslateResult struct {
	Source string `json:"src"`
	Target string `json:"tgt"`
}

type YoudaoTranslateResult struct {
	Type             string              `json:"type"`
	ErrorCode        int                 `json:"errorCode"`
	TranslateResults [][]TranslateResult `json:"translateResult"`
}

func SetupYoudaoTranslateCommand(rootCmd *cobra.Command) {
	var category string
	var list bool

	dict_support_category := map[string]string{
		"AUTO":     "自动检测语言",
		"ZH_CN2EN": "中文　»　英语",
		"ZH_CN2JA": "中文　»　日语",
		"ZH_CN2KR": "中文　»　韩语",
		"ZH_CN2FR": "中文　»　法语",
		"ZH_CN2RU": "中文　»　俄语",
		"ZH_CN2SP": "中文　»　西语",
		"ZH_CN2PT": "中文　»　葡语",
		"EN2ZH_CN": "英语　»　中文",
		"JA2ZH_CN": "日语　»　中文",
		"KR2ZH_CN": "韩语　»　中文",
		"FR2ZH_CN": "法语　»　中文",
		"RU2ZH_CN": "俄语　»　中文",
		"SP2ZH_CN": "西语　»　中文",
		"PT2ZH_CN": "葡语　»　中文",
	}
	cmd := &cobra.Command{
		Use:   "youdao-trans <word>",
		Short: "有道文章翻譯",
		Run: func(cmd *cobra.Command, args []string) {
			if list {
				var keys []string
				for k := range dict_support_category {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				for _, k := range keys {
					fmt.Printf("%s=%s\n", k, dict_support_category[k])
				}
				return
			}
			if len(args) < 1 {
				panic("required <word>")
			}
			word := args[0]
			category = strings.ToUpper(category)
			_, exists := dict_support_category[category]
			if !exists {
				category = "AUTO"
			}

			form := url.Values{}
			form.Add("type", category)
			form.Add("i", word)
			form.Add("xmlVersion", "1.8")
			form.Add("doctype", "json")
			form.Add("keyfrom", "fanyi.web")
			form.Add("ue", "UTF-8")
			form.Add("action", "FY_BY_CLICKBUTTON")
			form.Add("typoResult", "true")
			request := gorequest.New()
			_, body, err := request.Post("http://fanyi.youdao.com/translate").Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).
				Type("urlencoded").Send(form.Encode()).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			var result YoudaoTranslateResult
			json.Unmarshal([]byte(body), &result)
			// fmt.Println(result)
			if result.ErrorCode != 0 {
				fmt.Println("fail")
				return
			}
			// fmt.Println(result.Type)
			// fmt.Println(result)
			// fmt.Println(result.TranslateResults[0][0].Source)
			fmt.Println(result.TranslateResults[0][0].Target)
		},
	}

	cmd.Flags().StringVarP(&category, "category", "c", "AUTO", "翻譯類型")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "列出支援的翻譯類型")
	rootCmd.AddCommand(cmd)

}
