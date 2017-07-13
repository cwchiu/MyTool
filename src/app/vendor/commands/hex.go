package commands

import (
	"fmt"
    "encoding/hex"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func SetupHexCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "hex",
		Short: "hex view data",
		Long:  `hex view data`,
		Run: func(cmd *cobra.Command, args []string) {
            buf := make([]byte, 16)
            for {
                n, err := os.Stdin.Read(buf)
                if n == 0 || err == io.EOF {
                    break
                }
            
                if err != nil {
                    panic(err)
                }
                
                fmt.Println( hex.Dump( buf ) )
            }
		},
	}
	rootCmd.AddCommand(cmd)
}
