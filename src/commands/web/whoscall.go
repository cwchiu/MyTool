package web

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/cwchiu/MyTool/libs/api/whoscall"
)

func SetupWhosCallCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "whoscall <number>",
		Short: "查號碼來源",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("required <number>")
			}

			result, err := whoscall.Query(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result.Name)
				fmt.Println(result.Info)
			}
		},
	}
	rootCmd.AddCommand(cmd)

}
