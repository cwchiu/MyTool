package commands

import (
	ftp "commands/ftp"
	"github.com/spf13/cobra"
)

func SetupFtpCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "ftp", Short: "ftp"}

	ftp.SetupDownloadCommand(rootCmd)
	ftp.SetupUploadCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
