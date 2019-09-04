package imgur

import (
    "fmt"
	"github.com/spf13/cobra"
)

func SetupDeleteCommand(rootCmd *cobra.Command) {
	var cid string
	cmd := &cobra.Command{
		Use:   "delete <delete-hash>",
		Short: "刪除檔案",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}
            resp, err := DeleteImgur(cid, args[0])
            if err != nil {
                panic(err)
            }
            
            if resp.Success {
                fmt.Println("Ok")
            } else {
                err := resp.Data.(map[string]interface {})
                fmt.Printf("Status: %d, %s\n", resp.Status, err["error"])
            }
		},
	}

	cmd.Flags().StringVarP(&cid, "cid", "c", "", "Client ID")
	rootCmd.AddCommand(cmd)
}
