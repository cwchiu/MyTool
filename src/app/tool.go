package main

import (
	"github.com/spf13/cobra"
    commands "commands"
    "fmt"
)

// go build -o tool.exe
func main() {
    defer func() {
		if err := recover(); err != nil {
			fmt.Println( err.(string) )
		}
	}()
	rootCmd := &cobra.Command{Use: "tool"}
    commands.SetupFsCommand( rootCmd )
    commands.SetupWebCommand( rootCmd )
    commands.SetupServerCommand( rootCmd )
    commands.SetupHashCommand( rootCmd )
    commands.SetupMD2HtmlCommand( rootCmd )
    commands.SetupDateCommand( rootCmd )
	rootCmd.Execute()
}
