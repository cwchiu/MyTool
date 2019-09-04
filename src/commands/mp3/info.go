package mp3

import (
	"fmt"
	"os"

	"github.com/bogem/id3v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/spf13/cobra"
)

// SetupInfoCommand ...
func SetupInfoCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "info <filename>",
		Short: "mp3 資訊",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filename>")
			}

			f, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			defer f.Close()

			d, err := mp3.NewDecoder(f)
			if err != nil {
				panic(err)
			}
			// defer d.Close()
			fmt.Printf("SampleRate: %d\n", d.SampleRate())
			fmt.Printf("Length: %d Bytes\n", d.Length())

			tag, err := id3v2.Open(args[0], id3v2.Options{Parse: true})
			if err != nil {
				panic(err)
			}
			defer tag.Close()
			fmt.Printf("Artist: %s\n", tag.Artist())
			fmt.Printf("Album: %s\n", tag.Album())
			fmt.Printf("Title: %s\n", tag.Title())
			fmt.Printf("Genre: %s\n", tag.Genre())
			fmt.Printf("Year: %s\n", tag.Year())
			fmt.Printf("Size: %d Bytes\n", tag.Size())
		},
	}
	rootCmd.AddCommand(cmd)
}
