package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type result struct {
	Ip string `json:"ip"`
}

func ipify() string {
	request := gorequest.New()
	_, body, err := request.Get("https://api.ipify.org?format=json").End()
	if err != nil {
		panic(err)
	}

	var ret result
	err2 := json.Unmarshal([]byte(body), &ret)
	if err2 != nil {
		panic(err2)
	}
	return ret.Ip
}

func dnsomatic() string {
	request := gorequest.New()
	_, body, err := request.Get("http://myip.dnsomatic.com/").End()
	if err != nil {
		panic(err)
	}
	return body
}

func SetupMyipCommand(rootCmd *cobra.Command) {
	var src string
	cmd := &cobra.Command{
		Use:   "myip",
		Short: "取得外部IP",
		Run: func(cmd *cobra.Command, args []string) {
			if src == "dnsomatic" {
				fmt.Println(dnsomatic())
			} else {
				fmt.Println(ipify())
			}
		},
	}
	cmd.Flags().StringVarP(&src, "src", "s", "ipify", "服務供應商: ipify, dnsomatic, ")
	rootCmd.AddCommand(cmd)
}
