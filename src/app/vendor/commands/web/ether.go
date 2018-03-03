package web

import (
    api "libs/api/ether"
	"github.com/spf13/cobra"
)

func SetupEtherCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "ether",
		Short: "目前以太幣價格",
		Run: func(cmd *cobra.Command, args []string) {
			api.Query()
		},
	}
	rootCmd.AddCommand(cmd)

}
