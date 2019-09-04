package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func SetupPrettyCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "pretty",
		Short: "json pretty stdin",
		Long:  `json pretty stdin`,
		Run: func(cmd *cobra.Command, args []string) {
			var prettyJSON bytes.Buffer
			body, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			err = json.Indent(&prettyJSON, []byte(body), "", "\t")
			if err != nil {
				panic(err)
			}

			fmt.Println(string(prettyJSON.Bytes()))
		},
	}
	rootCmd.AddCommand(cmd)
}
