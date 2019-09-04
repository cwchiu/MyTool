package clamav

import (
	"github.com/spf13/cobra"
)

func SetupStatsCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Get Clamscan stats",
		Run:   createNoArgsCommandHandler("nSTATS"),
	}
	rootCmd.AddCommand(cmd)
}
