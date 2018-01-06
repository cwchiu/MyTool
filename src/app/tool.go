package main

import (
	commands "commands"
	arachni "commands/arachni"
	clamav "commands/clamav"
	"fmt"
	"github.com/spf13/cobra"
    _ "statik" 
)

// go build -o tool.exe
func main() {
	defer func() {
		if err := recover(); err != nil {
        fmt.Println(err.(string))
		}
	}()
	rootCmd := &cobra.Command{Use: "tool", Long: `我的個人常用工具
Site: https://chuiwenchiu.wordpress.com
Github: https://github.com/cwchiu/MyTool
    `}
	rootCmd.AddCommand(&cobra.Command{Use: "version", Short: "版本資訊", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.0.2")
	}})
	commands.SetupFsCommand(rootCmd)
	commands.SetupWebCommand(rootCmd)
	commands.SetupServerCommand(rootCmd)
	commands.SetupHashCommand(rootCmd)
	commands.SetupMD2HtmlCommand(rootCmd)
	commands.SetupDateCommand(rootCmd)
	commands.SetupBarCodeCommand(rootCmd)
	commands.SetupBase64Command(rootCmd)
	commands.SetupHexCommand(rootCmd)
	commands.SetupJsonCommand(rootCmd)
	commands.SetupGuerrillamailCommand(rootCmd)
	commands.SetupUrlCommand(rootCmd)
	commands.SetupPidCommand(rootCmd)
	commands.SetupPinyinCommand(rootCmd)
	commands.SetupNetcatCommand(rootCmd)
	commands.SetupJsCommand(rootCmd)
	commands.SetupMp3Command(rootCmd)
	commands.SetupWindowsCommand(rootCmd)
	commands.SetupImageCommand(rootCmd)
	commands.SetupImgurCommand(rootCmd)
	commands.SetupEpubCommand(rootCmd)
	commands.SetupCryptoCommand(rootCmd)
	commands.SetupSSHCommand(rootCmd)
	commands.SetupClipCommand(rootCmd)
	commands.SetupFtpCommand(rootCmd)
	commands.SetupSubtitleCommand(rootCmd)
	commands.SetupProxyPoolCommand(rootCmd)

    arachni.SetupArachniCommand(rootCmd)
    clamav.SetupCommand(rootCmd)
	rootCmd.Execute()
}
