package sqlite

import (
	"fmt"
	"github.com/spf13/cobra"
	lib "github.com/cwchiu/MyTool/libs/sqlite"
)


func setupQueryCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "query <addr> <key>",
		Short: "查詢",
        Long: `列出所有 Table : SELECT name FROM sqlite_master`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("<database> <sql>")
			}
			err := lib.Query(args[0], args[1], func (result map[string]interface{}) error{
                // fmt.Println(result)
                for k,v := range result {
                    if v == nil {
                        fmt.Printf("%s : NULL", k)
                        continue
                    }
                  switch v.(type) {
                    case int, int64:
                        fmt.Printf("%s : %d\n", k, v)
                    default:
                        fmt.Printf("%s : %v\n", k, string(v.([]byte)))
                  }
                }
                fmt.Println("-=-=-=-=-=-=-=-=-=-=-=")
                return nil
            })
			if err != nil {
				fmt.Println(err)
				return
			}

		},
	}
	rootCmd.AddCommand(cmd)
}

func SetupCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "sqlite", Short: "sqlite"}

	setupQueryCommand(cmd)

	rootCmd.AddCommand(cmd)
}
