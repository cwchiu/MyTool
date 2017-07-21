package main

import (
	commands "commands"
	"fmt"
	"github.com/spf13/cobra"
)

// go build -o tool.exe
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	rootCmd := &cobra.Command{Use: "tool", Long: `我的個人常用工具
    * Site: https://chuiwenchiu.wordpress.com
    * Github: https://github.com/cwchiu/MyTool
    `}
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
	commands.SetupTaiwanPidCommand(rootCmd)
	commands.SetupChinaPidCommand(rootCmd)
	commands.SetupPinyinCommand(rootCmd)
	commands.SetupNetcatCommand(rootCmd)
	commands.SetupJsCommand(rootCmd)
	commands.SetupMp3Command(rootCmd)
	commands.SetupWindowsCommand(rootCmd)

	rootCmd.Execute()
}
