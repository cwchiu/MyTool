package commands

import (
	image "github.com/cwchiu/MyTool/commands/image"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{Use: "image", Short: "影像"}

	image.SetupToPngCommand(cmd)
	image.SetupToGifCommand(cmd)
	image.SetupToJpegCommand(cmd)
	image.SetupToAsciiCommand(cmd)

	rootCmd.AddCommand(cmd)
}
