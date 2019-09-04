package clip

import (
	"github.com/atotto/clipboard"
	"fmt"
	"github.com/spf13/cobra"
)

func SetupGetCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "取得剪貼簿文字資料",
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := clipboard.ReadAll()
            if err != nil {
                panic(err)
            }
            fmt.Print(txt)
		},
	}
	rootCmd.AddCommand(cmd)
}
