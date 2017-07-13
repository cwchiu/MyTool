package commands

import (
	"github.com/spf13/cobra"
    fs "commands/fs"
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

    parentCmd.AddCommand(rootCmd)
	// rootCmd.Execute()
}
