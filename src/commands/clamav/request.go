package clamav

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetupRequestCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "request <command>",
		Short: "support PING,VERSION,RELOAD,SHUTDOWN",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <command>")
			}

			env, _ := GetClamAVEnviron()
			inst, err := CreateClamAV(env.Server)
			if err != nil {
				panic(err)
			}
			ver, err := inst.Request(args[0])
			if err != nil {
				panic(err)
			}
			fmt.Println(ver)

		},
	}
	rootCmd.AddCommand(cmd)
}
