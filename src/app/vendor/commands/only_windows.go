package commands

import (
	windows "commands/windows"
	"github.com/spf13/cobra"
)

func SetupWindowsCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "windows",
		Short: "Windows 相關",
	}

	windows.SetupLockCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
