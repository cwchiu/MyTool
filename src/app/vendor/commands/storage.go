package commands

import (
	bolt "commands/storage/bolt"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "storage", Short: "儲存引擎"}

	bolt.SetupCommand(cmd)

	rootCmd.AddCommand(cmd)
}
