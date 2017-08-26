package commands

import (
	clip "commands/clip"
	"github.com/spf13/cobra"
)

func SetupClipCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "clip",
		Short: "剪貼簿",
	}

	clip.SetupGetCommand(rootCmd)
	clip.SetupSetCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
