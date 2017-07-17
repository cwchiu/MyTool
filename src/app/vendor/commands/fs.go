package commands

import (
	fs "commands/fs"
	"github.com/spf13/cobra"
)

func SetupFsCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "fs"}

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

	parentCmd.AddCommand(rootCmd)
	// rootCmd.Execute()
}
