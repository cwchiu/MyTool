package crypto

import (
	"github.com/spf13/cobra"
)

func SetupRsaKeyDecryptCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "rsa-key-decrypt <file> <private key>",
		Short: "檔案解密",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("required <file> <private key>")
			}

			rsaKeyDecrypt(args[0], args[1])
		},
	}
	rootCmd.AddCommand(cmd)
}
