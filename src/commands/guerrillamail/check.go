package guerrillamail

import (
	"github.com/spf13/cobra"
)

func SetupCheckCommand(rootCmd *cobra.Command) {
	var token string
	cmd := &cobra.Command{
		Use:   "check",
		Short: "檢查是否有新郵件",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				panic("required -k <token>")
			}
			checkMail(token)
		},
	}
	cmd.Flags().StringVarP(&token, "token", "k", "", "Token")
	rootCmd.AddCommand(cmd)

}
