package proxypool

import (
	"github.com/spf13/cobra"
)

func setupCheckCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "初始化",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("TODO")
			// fmt.Println("check config.json")
			// fmt.Println("check pjs")
			check()
		},
	}
	rootCmd.AddCommand(cmd)
}
