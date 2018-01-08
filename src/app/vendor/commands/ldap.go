package commands

import (
    "commands/ldap"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "ldap", Short: "ldap api"}

	ldap.SetupSearchCommand(cmd)

	rootCmd.AddCommand(cmd)
}
