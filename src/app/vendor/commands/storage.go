package commands

import (
	bolt "commands/storage/bolt"
	redis "commands/storage/redis"
	sqlite "commands/storage/sqlite"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "storage", Short: "儲存引擎"}

	bolt.SetupCommand(cmd)
	redis.SetupCommand(cmd)
	sqlite.SetupCommand(cmd)

	rootCmd.AddCommand(cmd)
}
