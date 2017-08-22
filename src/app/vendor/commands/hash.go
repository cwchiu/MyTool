package commands

import (
	hash "commands/hash"
	"github.com/spf13/cobra"
)

func SetupHashCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "hash", Short: "雜湊計算"}

	hash.SetupMd5Command(rootCmd)
	hash.SetupSha1Command(rootCmd)
	hash.SetupSha256Command(rootCmd)
	hash.SetupSha384Command(rootCmd)
	hash.SetupSha512Command(rootCmd)
	hash.SetupCrc32Command(rootCmd)
	hash.SetupCertCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
