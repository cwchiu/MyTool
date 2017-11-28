package commands

import (
	"commands/web"
	"github.com/spf13/cobra"
)

func SetupWebCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "web",
		Short: "Web 服務相關",
	}

	web.SetupMyipCommand(rootCmd)
	web.SetupTlsVersionCommand(rootCmd)
	web.SetupUn53shareCommand(rootCmd)
	web.SetupDemd5Command(rootCmd)
	web.SetupDownloadCommand(rootCmd)
	web.SetupGoogleMapGeocodeCommand(rootCmd)
	web.SetupGoogleTranslateCommand(rootCmd)
	web.SetupIPInfoCommand(rootCmd)
	web.SetupUrlEncodeCommand(rootCmd)
	web.SetupUrlDecodeCommand(rootCmd)
	web.SetupExchangeRateCommand(rootCmd)
	// web.SetupYoudaoDictCommand(rootCmd)
	web.SetupGoogleDnsResolveCommand(rootCmd)
	web.SetupGenChineseNameCommand(rootCmd)
	web.SetupYoudaoTranslateCommand(rootCmd)
	web.SetupGistCommand(rootCmd)
	web.SetupSMSCommand(rootCmd)
	web.SetupMoreHandlinoCommand(rootCmd)
	web.SetupBabelGenCommand(rootCmd)
	web.SetupWhosCallCommand(rootCmd)
	web.SetupProxyCommand(rootCmd)
	web.SetupTinyCommand(rootCmd)
	web.SetupYandexTranslateCommand(rootCmd)
	web.SetupEtherCommand(rootCmd)
	web.SetupZipCodeCommand(rootCmd)
	web.SetupWeatherCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
