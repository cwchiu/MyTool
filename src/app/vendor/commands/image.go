package commands

import (
	image "commands/image"
	"github.com/spf13/cobra"
)

func SetupImageCommand(parentCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "image", Short: "影像"}

	image.SetupToPngCommand(cmd)
	image.SetupToGifCommand(cmd)
	image.SetupToJpegCommand(cmd)
	image.SetupToAsciiCommand(cmd)

	parentCmd.AddCommand(cmd)
}
