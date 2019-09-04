package crypto

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetupAesEncryptCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "aes-encrypt <file> <key>",
		Short: "檔案加密",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <file> <private key>")
			}
			key := rands(16)
			if len(args) > 1 {
				key_len := len(args[1])
				if key_len == 16 || key_len == 32 {
					key = args[1]
				}
			}
			fmt.Printf(`Key: %s`, key)
			aesEncrypt(args[0], key)
		},
	}
	rootCmd.AddCommand(cmd)
}
