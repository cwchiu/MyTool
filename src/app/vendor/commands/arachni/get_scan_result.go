package arachni

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

func SetupScanGetCommand(rootCmd *cobra.Command) {
	var username string
	var password string
	var server string

	cmd := &cobra.Command{
		Use:   "scan-get <task-id>",
		Short: "取得掃描結果",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <task-id>")
			}

			if server == "" {
				panic("required -s <arachni server host> ")
			}

			inst := CreateArachni(server, username, password)
			ret, err := inst.GetScanResult(args[0])
			if err != nil {
				panic(err)
			}
			// fmt.Println((*data)["seed"].(string))
			// fmt.Println(ret)

			// inst.StartScan(args[0], nil)
			// if  err != nil {
			// panic(err)
			// }
			bs, err := json.MarshalIndent(ret, "", "    ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
			// fmt.Println("------------------")
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "arachni server username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "arachni server password")
	cmd.Flags().StringVarP(&server, "server", "s", "", "arachni server host")
	rootCmd.AddCommand(cmd)
}
