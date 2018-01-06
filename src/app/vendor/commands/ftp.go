package commands

import (
	ftp "commands/ftp"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "ftp", Short: "ftp"}

	ftp.SetupDownloadCommand(cmd)
	ftp.SetupUploadCommand(cmd)

	rootCmd.AddCommand(cmd)
}
