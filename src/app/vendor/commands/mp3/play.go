package netcat

import (
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/spf13/cobra"
	"io"
    "bufio"
    "strings"
	"os"
    "fmt"
)

// https://github.com/hajimehoshi/go-mp3/blob/master/example/main.go

func play(fn string) {
    f, err := os.Open(fn)
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


    if _, err := io.Copy(p, d); err != nil {
        panic(err)
    }
}
func SetupPlayCommand(rootCmd *cobra.Command) {
	var repeat int
	cmd := &cobra.Command{
		Use:   "play <filename>",
		Short: "播放mp3",
		Run: func(cmd *cobra.Command, args []string) {
            var playlist []string
			if len(args) < 1 {
				reader := bufio.NewReader( os.Stdin )
                
                for {
                    line, err := reader.ReadString('\n')
                    if len(line) > 0 {
                        // fmt.Println(line)
                        playlist = append(playlist, strings.TrimSpace(line))
                    }
                    
                    if err == io.EOF {
                        break
                    }
                    
                    if err != nil {
                        panic(err)
                    }
                }
			}else{
                playlist = args
            }
            fmt.Printf("Music: %d\n", len(playlist) )
            c := 1
            for {
                for _, fn := range playlist {
                    fmt.Printf("Playing [%s]\n", fn)
                    play(fn)
                }
                
                if repeat > 0 && c >= repeat {
                    break
                }
                c += 1
            }
		},
	}

	cmd.Flags().IntVarP(&repeat, "repeat", "r", 1, "重複次數, 小於 1 表示無限重複")
	rootCmd.AddCommand(cmd)
}
