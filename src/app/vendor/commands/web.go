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
	web.SetupDownloadCommand(rootCmd)
	web.SetupGoogleMapGeocodeCommand(rootCmd)
	web.SetupIPInfoCommand(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
