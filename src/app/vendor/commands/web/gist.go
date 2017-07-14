package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type GistFileContent struct {
	Content string `json:"content"`
}

type GistForm struct {
	Desc   string                     `json:"description"`
	Public bool                       `json:"public"`
	Files  map[string]GistFileContent `json:"files"`
}

type GistResponse struct {
	Url string `json:"html_url"`
}

func SetupGistCommand(rootCmd *cobra.Command) {
	var title string
	var desc string
	var public bool
	var username string
	var token string

	cmd := &cobra.Command{
		Use:   "gist <word>",
		Short: "新增 gist",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <word>")
			}

			data := GistForm{
				Desc:   desc,
				Public: public,
				Files: map[string]GistFileContent{
					title: GistFileContent{
						Content: args[0],
					},
				},
			}

			request := gorequest.New().Post("https://api.github.com/gists").SetDebug(false).Send(data)

			if username != "" && token != "" {
				request.Set("User-Agent", "Awesome-Octocat-App")
				hash := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
				request.Set("Authorization", "Basic "+hash)
			} else {
				request.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")
			}
			_, body, err := request.End()
			if err != nil {
				panic(err)
			}

			// fmt.Println(body)
			var result GistResponse
			json.Unmarshal([]byte(body), &result)
			fmt.Println(result.Url)
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "帳號")
	cmd.Flags().StringVarP(&token, "token", "k", "", "密碼")
	cmd.Flags().StringVarP(&title, "title", "t", "Unknown.txt", "標題")
	cmd.Flags().StringVarP(&desc, "desc", "d", "the description for this gist", "描述")
	cmd.Flags().BoolVarP(&public, "public", "p", false, "公開?")
	rootCmd.AddCommand(cmd)

}
