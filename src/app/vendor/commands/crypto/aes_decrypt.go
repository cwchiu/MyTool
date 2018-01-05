package crypto

import (
	"github.com/spf13/cobra"
)

    
    
func SetupAesDecryptCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "aes-decrypt <file> <private key>",
		Short: "檔案解密",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 2 {
				panic("required <file> <private key>")
			}
            key_len := len(args[1])
            if key_len != 16 && key_len != 32 {
                panic("key length must 16 or 32")
            }
            aesDecrypt(args[0], args[1])
		},
	}
	rootCmd.AddCommand(cmd)
}
