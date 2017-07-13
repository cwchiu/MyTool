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
    SetupFsCommand( rootCmd )
    SetupWebCommand( rootCmd )
    SetupServerCommand( rootCmd )
    commands.SetupMD2HtmlCommand( rootCmd )
	rootCmd.Execute()
}
