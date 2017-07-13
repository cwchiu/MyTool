package commands

import (
	"github.com/spf13/cobra"
    "commands/web"
)

func SetupWebCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "web"}

	web.SetupMyipCommand(rootCmd)
	web.SetupTlsVersionCommand(rootCmd)
	web.SetupUn53shareCommand(rootCmd)
	web.SetupDemd5Command(rootCmd)

	// rootCmd.Execute()
    parentCmd.AddCommand(rootCmd)
}
