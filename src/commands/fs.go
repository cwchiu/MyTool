package commands

import (
	fs "github.com/cwchiu/MyTool/commands/fs"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "fs", Short: "檔案相關操作"}

	fs.SetupEmptyCommand(cmd)
	fs.SetupInfoCommand(cmd)
	fs.SetupTruncatCommand(cmd)
	fs.SetupCatCommand(cmd)
	fs.SetupMvCommand(cmd)
	fs.SetupRmCommand(cmd)
	fs.SetupZipCommand(cmd)
	fs.SetupUnzipCommand(cmd)
	fs.SetupLsCommand(cmd)
	fs.SetupTreeCommand(cmd)
	fs.SetupCpCommand(cmd)
	fs.SetupNlCommand(cmd)
	fs.SetupDos2UnixCommand(cmd)
	fs.SetupUnix2DosCommand(cmd)
	fs.SetupCutCommand(cmd)
	fs.SetupGrepCommand(cmd)

	rootCmd.AddCommand(cmd)
}
