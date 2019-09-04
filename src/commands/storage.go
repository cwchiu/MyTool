package commands

import (
	bolt "github.com/cwchiu/MyTool/commands/storage/bolt"
	redis "github.com/cwchiu/MyTool/commands/storage/redis"
	// sqlite "github.com/cwchiu/MyTool/commands/storage/sqlite"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "storage", Short: "儲存引擎"}

	bolt.SetupCommand(cmd)
	redis.SetupCommand(cmd)
	// sqlite.SetupCommand(cmd)

	rootCmd.AddCommand(cmd)
}
