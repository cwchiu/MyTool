package clip

import (
	"github.com/atotto/clipboard"
	"fmt"
	"github.com/spf13/cobra"
    "io/ioutil"
    "os"
)

func SetupSetCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "stdin 寫入剪貼簿",
		Run: func(cmd *cobra.Command, args []string) {
            bs, err := ioutil.ReadAll(os.Stdin)
            if err != nil {
                panic(err)
            }
            
			err = clipboard.WriteAll(string(bs))
            if err != nil {
                panic(err)
            }
            fmt.Println("Ok")
		},
	}
	rootCmd.AddCommand(cmd)
}
