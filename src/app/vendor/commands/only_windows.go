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
	windows.SetupWord2TxtCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
