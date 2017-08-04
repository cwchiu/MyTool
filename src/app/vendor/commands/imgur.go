package commands

import (
	"commands/imgur"
	"github.com/spf13/cobra"
)

func SetupImgurCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "imgur",
		Short: "imgur 服務相關",
	}

	imgur.SetupUploadCommand(rootCmd)
	imgur.SetupDeleteCommand(rootCmd)


	parentCmd.AddCommand(rootCmd)
}
