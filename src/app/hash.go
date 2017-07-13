package main

import (
	"github.com/spf13/cobra"
    hash "commands/hash"
)


func SetupHashCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{Use: "hash"}

	hash.SetupMd5Command(rootCmd)
	hash.SetupSha1Command(rootCmd)
	hash.SetupSha256Command(rootCmd)
	hash.SetupSha384Command(rootCmd)
	hash.SetupSha512Command(rootCmd)
	hash.SetupCrc32Command(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
