package web

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

func SetupTinyCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "tiny <url>",
		Short: "tinyurl",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <url>")
			}
			request := gorequest.New()
			_, body, err := request.Get("http://tinyurl.com/api-create.php").Param("url", args[0]).End()
			if err != nil {
				panic(err)
			}
			fmt.Println(body)

		},
	}

	rootCmd.AddCommand(cmd)

}
