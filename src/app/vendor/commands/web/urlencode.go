package web

import (
	"fmt"
	"net/url"
	"github.com/spf13/cobra"
)


func SetupUrlEncodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "urlencode <code>",
		Short: "string urlencode",
		Long:  `string urlencode`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("required <ip>")
			}
            
            fmt.Println(url.QueryEscape(args[0]))
		},
	}
	rootCmd.AddCommand(cmd)

}
