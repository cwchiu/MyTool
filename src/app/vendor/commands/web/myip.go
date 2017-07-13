package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type Result struct {
	Ip string `json:"ip"`
}

func SetupMyipCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "myip",
		Short: "myip",
		Long:  `myip`,
		Run: func(cmd *cobra.Command, args []string) {

			request := gorequest.New()
			_, body, err := request.Get("https://api.ipify.org?format=json").End()
			if err != nil {
				panic(err)
			}

			var ret Result
			err2 := json.Unmarshal([]byte(body), &ret)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println(ret.Ip)

		},
	}

	rootCmd.AddCommand(cmd)

}
