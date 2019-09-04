package commands

import (
	windows "github.com/cwchiu/MyTool/commands/windows"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "windows",
		Short: "Windows 相關",
	}

	windows.SetupLockCommand(cmd)
	windows.SetupWord2TxtCommand(cmd)

	rootCmd.AddCommand(cmd)
}
