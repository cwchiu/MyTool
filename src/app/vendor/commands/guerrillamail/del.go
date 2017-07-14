package guerrillamail

import (
	"github.com/spf13/cobra"
)

func SetupDelCommand(rootCmd *cobra.Command) {
	var token string
	var email_id string
	cmd := &cobra.Command{
		Use:   "del",
		Short: "刪除郵件",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				panic("required -k <token>")
			}
			if email_id == "" {
				panic("required -i <email_id>")
			}
			delMail(token, email_id)
			// email, err := mailClient.SetEmailUser(guerrillamail.Argument{
			// "email_user": "qianlnk",
			// "lang": guerrillamail.LANGUAGE_ZH_HANT,
			// })

		},
	}
	cmd.Flags().StringVarP(&token, "token", "k", "", "Token")
	cmd.Flags().StringVarP(&email_id, "id", "i", "", "Email ID")
	rootCmd.AddCommand(cmd)

}
