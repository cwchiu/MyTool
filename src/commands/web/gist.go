package web

import (

	"fmt"
    "github.com/cwchiu/MyTool/libs/api/github"
	"github.com/spf13/cobra"
)


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
            cfg := github.NewGistConfig()
            cfg.Title = title
            cfg.Content = args[0]
            cfg.Username = username
            cfg.Token = token
            cfg.Desc = desc
            resp, err := github.CreateGist( cfg )
            if err != nil {
                fmt.Println(err)
            }else{
                fmt.Println(resp.Url)
            }
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "帳號")
	cmd.Flags().StringVarP(&token, "token", "k", "", "密碼")
	cmd.Flags().StringVarP(&title, "title", "t", "Unknown.txt", "標題")
	cmd.Flags().StringVarP(&desc, "desc", "d", "the description for this gist", "描述")
	cmd.Flags().BoolVarP(&public, "public", "p", false, "公開?")
	rootCmd.AddCommand(cmd)

}
