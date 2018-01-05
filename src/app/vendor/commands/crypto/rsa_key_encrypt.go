package crypto

import (
	"github.com/spf13/cobra"
)
    
func SetupRsaKeyEncryptCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "rsa-key-encrypt <file> <public key>",
		Short: "檔案加密",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 2 {
				panic("required <file> <public key>")
			}
            
            rsaKeyEncrypt(args[0], args[1])
		},
	}
	rootCmd.AddCommand(cmd)
}
