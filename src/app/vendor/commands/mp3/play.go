package netcat

import (
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// https://github.com/hajimehoshi/go-mp3/blob/master/example/main.go
func SetupPlayCommand(rootCmd *cobra.Command) {
	var repeat int
	cmd := &cobra.Command{
		Use:   "play <filename>",
		Short: "播放mp3",
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
			defer d.Close()

			p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
			if err != nil {
				panic(err)
			}
			defer p.Close()

			c := 0
			for {
				if _, err := io.Copy(p, d); err != nil {
					panic(err)
				}
				c += 1
				if repeat > 0 && c >= repeat {
					break
				}

				d.Seek(0, io.SeekStart)
			}
		},
	}

	cmd.Flags().IntVarP(&repeat, "repeat", "r", 1, "重複次數, 小於 1 表示無限重複")
	rootCmd.AddCommand(cmd)
}
