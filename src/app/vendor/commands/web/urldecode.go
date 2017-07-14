package web

import (
	"fmt"
	"net/url"
	"github.com/spf13/cobra"
)


func SetupUrlDecodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "urldecode <code>",
		Short: "string urldecode",
		Long:  `string urldecode`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("required <ip>")
			}
            r, err := url.QueryUnescape(args[0])
            if err != nil {
                panic(err)
            }
            fmt.Println(r)
		},
	}
	rootCmd.AddCommand(cmd)

}
