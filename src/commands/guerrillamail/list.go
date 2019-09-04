package guerrillamail

import (
	"github.com/spf13/cobra"
)

func SetupListCommand(rootCmd *cobra.Command) {
	var token string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出所有郵件",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				panic("required -k <token>")
			}
			getEmailList(token)
		},
	}
	cmd.Flags().StringVarP(&token, "token", "k", "", "Token")
	rootCmd.AddCommand(cmd)
}
