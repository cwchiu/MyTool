package crypto

import (
	"github.com/spf13/cobra"
)
    
func SetupAesEncryptCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "aes-encrypt <file> <key>",
		Short: "檔案加密",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 2 {
				panic("required <file> <key>")
			}
            
            key_len := len(args[1])
            if key_len != 16 && key_len != 32 {
                panic("key length must 16 or 32")
            }
            aesEncrypt(args[0], args[1])
		},
	}
	rootCmd.AddCommand(cmd)
}
