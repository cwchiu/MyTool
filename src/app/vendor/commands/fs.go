package commands

import (
	fs "commands/fs"
	"github.com/spf13/cobra"
)

func SetupFsCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "fs", Short: "檔案相關操作"}

	fs.SetupEmptyCommand(rootCmd)
	fs.SetupInfoCommand(rootCmd)
	fs.SetupTruncatCommand(rootCmd)
	fs.SetupCatCommand(rootCmd)
	fs.SetupMvCommand(rootCmd)
	fs.SetupRmCommand(rootCmd)
	fs.SetupZipCommand(rootCmd)
	fs.SetupUnzipCommand(rootCmd)
	fs.SetupLsCommand(rootCmd)
	fs.SetupTreeCommand(rootCmd)
	fs.SetupCpCommand(rootCmd)
	fs.SetupNlCommand(rootCmd)
	fs.SetupDos2UnixCommand(rootCmd)
	fs.SetupUnix2DosCommand(rootCmd)
	fs.SetupCutCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
	// rootCmd.Execute()
}
