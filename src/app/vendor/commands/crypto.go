package commands

import (
	crypto "commands/crypto"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "crypto", Short: "加密"}

	crypto.SetupCertCommand(cmd)
	crypto.SetupApkCertCommand(cmd)
	crypto.SetupSignCommand(cmd)
	crypto.SetupVerifyCommand(cmd)
	crypto.SetupGenRsaKeyCommand(cmd)
	crypto.SetupRsaKeyEncryptCommand(cmd)
	crypto.SetupRsaKeyDecryptCommand(cmd)
	crypto.SetupAesEncryptCommand(cmd)
	crypto.SetupAesDecryptCommand(cmd)

	rootCmd.AddCommand(cmd)
}
