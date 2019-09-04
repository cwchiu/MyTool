package commands

import (
	"github.com/cwchiu/MyTool/commands/web"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "web",
		Short: "Web 服務相關",
	}

	web.SetupMyipCommand(cmd)
	web.SetupTlsVersionCommand(cmd)
	web.SetupUn53shareCommand(cmd)
	web.SetupDemd5Command(cmd)
	web.SetupDownloadCommand(cmd)
	web.SetupGoogleMapGeocodeCommand(cmd)
	web.SetupGoogleTranslateCommand(cmd)
	web.SetupIPInfoCommand(cmd)
	web.SetupUrlEncodeCommand(cmd)
	web.SetupUrlDecodeCommand(cmd)
	web.SetupExchangeRateCommand(cmd)
	// web.SetupYoudaoDictCommand(cmd)
	web.SetupGoogleDnsResolveCommand(cmd)
	web.SetupGenChineseNameCommand(cmd)
	web.SetupGistCommand(cmd)
	web.SetupSMSCommand(cmd)
	web.SetupMoreHandlinoCommand(cmd)
	web.SetupBabelGenCommand(cmd)
	web.SetupWhosCallCommand(cmd)
	web.SetupProxyCommand(cmd)
	web.SetupTinyCommand(cmd)
	web.SetupYoudaoCommand(cmd)
	web.SetupEtherCommand(cmd)
	web.SetupZipCodeCommand(cmd)
	web.SetupWeatherCommand(cmd)
	web.SetupPhoneCommand(cmd)
	web.SetupPasteBinCommand(cmd)
	web.SetupMusic163Command(cmd)

	rootCmd.AddCommand(cmd)
}
