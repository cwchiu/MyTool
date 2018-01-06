package commands

import (
	"github.com/spf13/cobra"
    "commands/clamav"
)

func init() {
	cmd := &cobra.Command{Use: "clamav", Short: "ClamAV api"}

	clamav.SetupRequestCommand(cmd)
    
	clamav.SetupVersionCommand(cmd)
	clamav.SetupPingCommand(cmd)
	clamav.SetupReloadCommand(cmd)
	clamav.SetupShutdownCommand(cmd)
	clamav.SetupStatsCommand(cmd)
	clamav.SetupCommandListCommand(cmd)
    
	clamav.SetupScanCommand(cmd)
	clamav.SetupAllMatchScanCommand(cmd)
	clamav.SetupContScanCommand(cmd)
	clamav.SetupMultiScanCommand(cmd)

	rootCmd.AddCommand(cmd)
}
