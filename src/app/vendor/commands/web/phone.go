package web

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type infoData struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Sp       string `json:"sp"`
}
type phoneInfo struct {
	Code int      `json:"code"`
	Data infoData `json:"data"`
}

func SetupPhoneCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "phone <number>",
		Short: "根據360的api獲取到某電話的具體歸屬地",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("required <number>")
			}
			request := getRequest(false, false)
			_, body, err := request.Get(fmt.Sprintf("http://cx.shouji.360.cn/phonearea.php?number=%s", args[0])).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			var result phoneInfo
			json.Unmarshal([]byte(body), &result)
			fmt.Println(result.Data.Province)
			fmt.Println(result.Data.Sp)
			fmt.Println(result.Data.City)
		},
	}
	rootCmd.AddCommand(cmd)

}
