package commands

import (
	"github.com/spf13/cobra"
    "commands/web"
)

func SetupWebCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
        Use: "web",
        Short: "Web 服務相關",
    }

	web.SetupMyipCommand(rootCmd)
	web.SetupTlsVersionCommand(rootCmd)
	web.SetupUn53shareCommand(rootCmd)
	web.SetupDemd5Command(rootCmd)
	web.SetupDownloadCommand(rootCmd)
	web.SetupGoogleMapGeocodeCommand(rootCmd)
	web.SetupIPInfoCommand(rootCmd)
	web.SetupUrlEncodeCommand(rootCmd)
	web.SetupUrlDecodeCommand(rootCmd)
	web.SetupExchangeRateCommand(rootCmd)
	web.SetupYoudaoDictCommand(rootCmd)
	web.SetupGoogleDnsResolveCommand(rootCmd)
	web.SetupGenChineseNameCommand(rootCmd)
	web.SetupYoudaoTranslateCommand(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
