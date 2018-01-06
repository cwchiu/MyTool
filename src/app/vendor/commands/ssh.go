package commands

import (
	ssh "commands/ssh"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "ssh", Short: "ssh"}

	ssh.SetupExecCommand(cmd)
	ssh.SetupUploadCommand(cmd)
	ssh.SetupDownloadCommand(cmd)
	ssh.SetupTtyCommand(cmd)

	rootCmd.AddCommand(cmd)
}
