package web

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

func SetupDemd5Command(rootCmd *cobra.Command) {
    var uid string
	var token string
	cmd := &cobra.Command{
		Use:   "demd5 <code>",
		Short: "md5 decode",
		Long:  `md5 decode`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("need an code")
			}
			request := gorequest.New()
            code := args[0]
			_, body, err := request.Get("http://www.ttmd5.com/do.php").Param("c","Api").Param("m", "crack").Param("uid", uid).Param("token", token).Param("cipher", code).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			fmt.Println(body)
		},
	}
	cmd.Flags().StringVarP(&uid, "uid", "u", "", "uid")
	cmd.Flags().StringVarP(&token, "token", "t", "", "token")
	rootCmd.AddCommand(cmd)

}
