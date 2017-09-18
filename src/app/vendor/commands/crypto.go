package commands

import (
	crypto "commands/crypto"
	"github.com/spf13/cobra"
)

func SetupCryptoCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "crypto", Short: "加密"}

	crypto.SetupCertCommand(rootCmd)
	crypto.SetupApkCertCommand(rootCmd)
	crypto.SetupSignCommand(rootCmd)
	crypto.SetupVerifyCommand(rootCmd)
	crypto.SetupGenRsaKeyCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
