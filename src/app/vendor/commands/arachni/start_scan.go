package arachni

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

func SetupScanStartCommand(rootCmd *cobra.Command) {
	var username string
	var password string
	var server string

	cmd := &cobra.Command{
		Use:   "scan-start <url>",
		Short: "啟動掃描",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <url>")
			}

			if server == "" {
				panic("required -s <arachni server host> ")
			}

			inst := CreateArachni(server, username, password)
			// "http://192.168.99.100:7331"
			// m := map[string]interface{}{
			// "checks": []string{"xss*", "sql*"},
			// "audit": map[string]interface{}{
			// "elements": []string{"links", "forms"},
			// },
			// }
			// ret, err := inst.StartScan("http://demo.testfire.net/default.aspx", &m)
			ret, err := inst.StartScan(args[0], nil)
			if err != nil {
				panic(err)
			}
			// fmt.Println(json.MarshalIndent(ret, "", "    "))
			bs, err := json.MarshalIndent(ret, "", "    ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "arachni server username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "arachni server password")
	cmd.Flags().StringVarP(&server, "server", "s", "", "arachni server host")
	rootCmd.AddCommand(cmd)
}
