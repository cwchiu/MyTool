package web

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/cwchiu/MyTool/libs/api/youdao"
	// "sort"
	// "strings"
)

func setupYoudaoDictCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "dict <word>",
		Short: "字典詞語翻譯",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <word>")
			}
			result, err := youdao.DictQuery(args[0])
			if err == nil {
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		},
	}
	rootCmd.AddCommand(cmd)

}

func SetupYoudaoCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "youdao",
		Short: "有道服務",
	}

	setupYoudaoTranslateCommand(cmd)
	setupYoudaoDictCommand(cmd)

	rootCmd.AddCommand(cmd)
}

func setupYoudaoTranslateCommand(rootCmd *cobra.Command) {
	var from string
	var to string
	// var list bool

	// dict_support_category := map[string]string{
	// "AUTO":     "自动检测语言",
	// "ZH_CN2EN": "中文　»　英语",
	// "ZH_CN2JA": "中文　»　日语",
	// "ZH_CN2KR": "中文　»　韩语",
	// "ZH_CN2FR": "中文　»　法语",
	// "ZH_CN2RU": "中文　»　俄语",
	// "ZH_CN2SP": "中文　»　西语",
	// "ZH_CN2PT": "中文　»　葡语",
	// "EN2ZH_CN": "英语　»　中文",
	// "JA2ZH_CN": "日语　»　中文",
	// "KR2ZH_CN": "韩语　»　中文",
	// "FR2ZH_CN": "法语　»　中文",
	// "RU2ZH_CN": "俄语　»　中文",
	// "SP2ZH_CN": "西语　»　中文",
	// "PT2ZH_CN": "葡语　»　中文",
	// }
	cmd := &cobra.Command{
		Use:   "translate <word>",
		Short: "文章翻譯",
		Run: func(cmd *cobra.Command, args []string) {
			// if list {
			// var keys []string
			// for k := range dict_support_category {
			// keys = append(keys, k)
			// }
			// sort.Strings(keys)
			// for _, k := range keys {
			// fmt.Printf("%s=%s\n", k, dict_support_category[k])
			// }
			// return
			// }
			if len(args) < 1 {
				panic("required <word>")
			}
			word := args[0]
			// category = strings.ToUpper(category)
			// _, exists := dict_support_category[category]
			// if !exists {
			// category = "AUTO"
			// }

			result, err := youdao.Translate(word, "AUTO", "AUTO")

			if err == nil {
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		},
	}
	lang := "AUTO(自動識別), ZH_CN(中文), EN(英语) ,JA(日语), KR(韩语), FR(法语), RU(俄语), SP(西语), PT(葡语)"
	cmd.Flags().StringVarP(&from, "from", "f", "AUTO", "來源語言:"+lang)
	cmd.Flags().StringVarP(&to, "to", "t", "AUTO", "目的語言:"+lang)
	// cmd.Flags().BoolVarP(&list, "list", "l", false, "列出支援的翻譯類型")
	rootCmd.AddCommand(cmd)

}
