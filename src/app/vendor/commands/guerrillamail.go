package commands

import (
	"github.com/spf13/cobra"
    guerrillamail "commands/guerrillamail"
)

func SetupGuerrillamailCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
        Use: "guerrillamail",
        Short: "臨時用郵件",
    }

	guerrillamail.SetupNewCommand(rootCmd)
	guerrillamail.SetupListCommand(rootCmd)
	guerrillamail.SetupCheckCommand(rootCmd)
	guerrillamail.SetupFetchCommand(rootCmd)
	guerrillamail.SetupDelCommand(rootCmd)

    parentCmd.AddCommand(rootCmd)
}
