package commands

import (
	hash "github.com/cwchiu/MyTool/commands/hash"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "hash", Short: "雜湊計算"}

	hash.SetupMd5Command(cmd)
	hash.SetupSha1Command(cmd)
	hash.SetupSha256Command(cmd)
	hash.SetupSha384Command(cmd)
	hash.SetupSha512Command(cmd)
	hash.SetupCrc32Command(cmd)

	rootCmd.AddCommand(cmd)
}
