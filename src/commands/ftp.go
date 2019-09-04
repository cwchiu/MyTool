package commands

import (
	ftp "github.com/cwchiu/MyTool/commands/ftp"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "ftp", Short: "ftp"}

	ftp.SetupDownloadCommand(cmd)
	ftp.SetupUploadCommand(cmd)

	rootCmd.AddCommand(cmd)
}
