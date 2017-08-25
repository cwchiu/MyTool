package commands

import (
	ssh "commands/ssh"
	"github.com/spf13/cobra"
)

func SetupSSHCommand(parentCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "ssh", Short: "ssh"}

	ssh.SetupExecCommand(cmd)
	ssh.SetupUploadCommand(cmd)
	ssh.SetupDownloadCommand(cmd)
	ssh.SetupTtyCommand(cmd)

	parentCmd.AddCommand(cmd)
}
