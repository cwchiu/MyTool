package commands

import (
	"github.com/cwchiu/MyTool/commands/imgur"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "imgur",
		Short: "imgur 服務相關",
	}

	imgur.SetupUploadCommand(cmd)
	imgur.SetupDeleteCommand(cmd)


	rootCmd.AddCommand(cmd)
}
