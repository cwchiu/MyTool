package main

import (
	"github.com/spf13/cobra"
)

// go build -o tool.exe
func main() {
	rootCmd := &cobra.Command{Use: "tool"}
    SetupFsCommand( rootCmd )
    SetupWebCommand( rootCmd )
    SetupServerCommand( rootCmd )
	rootCmd.Execute()
}
