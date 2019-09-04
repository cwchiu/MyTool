package commands

import (
	clip "github.com/cwchiu/MyTool/commands/clip"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "clip",
		Short: "剪貼簿",
	}

	clip.SetupGetCommand(cmd)
	clip.SetupSetCommand(cmd)

	rootCmd.AddCommand(cmd)
}
