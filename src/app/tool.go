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
	rootCmd := &cobra.Command{Use: "tool"}
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

	rootCmd.Execute()
}
