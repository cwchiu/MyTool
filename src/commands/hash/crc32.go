package hash

import (
	"bufio"
	"hash/crc32"
	"fmt"
    "io"
	"github.com/spf13/cobra"
	"os"
    "strconv"
)

func SetupCrc32Command(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "crc32",
		Short: "crc32 stdin",
		Long:  `crc32 stdin`,
		Run: func(cmd *cobra.Command, args []string) {
            table := crc32.MakeTable(crc32.IEEE)
            var crc32_hash uint32 = 0
    
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
                
                if crc32_hash == 0 {
                    crc32_hash = crc32.Checksum(buf, table) 
                }else {
                    crc32_hash = crc32.Update(crc32_hash, table, buf)
                }
			}

            fmt.Println(strconv.FormatUint(uint64(crc32_hash), 16))
		},
	}
	rootCmd.AddCommand(cmd)
}
