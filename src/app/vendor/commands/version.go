package commands

import (
    "fmt"
	"github.com/spf13/cobra"
)

const APP_VERSION = "0.0.3"
func init() {
	rootCmd.AddCommand(&cobra.Command{Use: "version", Short: "版本資訊", Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", APP_VERSION)
	}})
}
