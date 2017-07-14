package guerrillamail

import (
	"github.com/spf13/cobra"
)

func SetupNewCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "取得一個新的臨時電子信箱",
		Run: func(cmd *cobra.Command, args []string) {
			getEmail()
		},
	}

	rootCmd.AddCommand(cmd)
}
