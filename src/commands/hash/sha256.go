package hash

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
    "io"
	"github.com/spf13/cobra"
	"os"
)

func SetupSha256Command(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "sha256",
		Short: "sha256 stdin",
		Long:  `sha256 stdin`,
		Run: func(cmd *cobra.Command, args []string) {
			h := sha256.New()

			r := bufio.NewReader(os.Stdin)
			buf := make([]byte, 0, 4*1024)
			for {
				n, err := r.Read(buf[:cap(buf)])
				buf = buf[:n]
				if n == 0 {
					if err == nil {
						continue
					}
					if err == io.EOF {
						break
					}
					panic(err)
				}
				if err != nil && err != io.EOF {
					panic(err)
				}

				h.Write(buf)
                
			}

			fmt.Println(hex.EncodeToString(h.Sum(nil)))
		},
	}
	rootCmd.AddCommand(cmd)
}
