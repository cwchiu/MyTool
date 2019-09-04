package commands

import (
	guerrillamail "github.com/cwchiu/MyTool/commands/guerrillamail"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "guerrillamail",
		Short: "臨時用郵件",
	}

	guerrillamail.SetupNewCommand(cmd)
	guerrillamail.SetupListCommand(cmd)
	guerrillamail.SetupCheckCommand(cmd)
	guerrillamail.SetupFetchCommand(cmd)
	guerrillamail.SetupDelCommand(cmd)

	rootCmd.AddCommand(cmd)
}
