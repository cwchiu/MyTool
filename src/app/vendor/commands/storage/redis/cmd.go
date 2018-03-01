package redis

import (
	"fmt"
	"github.com/spf13/cobra"
	lib "libs/redis"
)

func setupPutCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "put <redis> <key> <value>",
		Short: "儲存",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				panic("<addr> <key> <value>")
			}
			err := lib.Put(args[0], args[1], args[2])
			if err != nil {
				panic(err)
			}

			fmt.Println("Ok")
		},
	}
	rootCmd.AddCommand(cmd)
}

func setupGetCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "get <addr> <key>",
		Short: "取值",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("<addr> <key>")
			}
			value, err := lib.Get(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(value)
		},
	}
	rootCmd.AddCommand(cmd)
}

func SetupCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "redis", Short: "redis"}

	setupPutCommand(cmd)
	setupGetCommand(cmd)

	rootCmd.AddCommand(cmd)
}
