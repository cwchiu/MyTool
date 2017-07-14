package guerrillamail

import (
	"github.com/spf13/cobra"
)

func SetupFetchCommand(rootCmd *cobra.Command) {
	var token string
	var email_id string
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "取郵件內文",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				panic("required -k <token>")
			}
			if email_id == "" {
				panic("required -i <email_id>")
			}
			fetchMail(token, email_id)
		},
	}
	cmd.Flags().StringVarP(&token, "token", "k", "", "Token")
	cmd.Flags().StringVarP(&email_id, "id", "i", "", "Email ID")
	rootCmd.AddCommand(cmd)

}
