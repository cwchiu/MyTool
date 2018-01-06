package commands

import (
    "fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{Use: "version", Short: "版本資訊", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.0.2")
	}})
}
