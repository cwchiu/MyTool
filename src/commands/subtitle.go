package commands

import (
	subtitle "github.com/cwchiu/MyTool/commands/subtitle"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "subtitle", Short: "字幕"}

	subtitle.SetupAss2SrtCommand(cmd)
	subtitle.SetupSrt2AssCommand(cmd)

	rootCmd.AddCommand(cmd)
}
