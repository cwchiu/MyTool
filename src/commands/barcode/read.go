package barcode

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tuotoo/qrcode"
	"os"
)

// SetupReadCommand QRCode 掃描命令列入口
func SetupReadCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "read <text>",
		Short: "read qr code",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <text>")
			}
			fi, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			defer fi.Close()
            // qrcode.Debug = true
			qrmatrix, err := qrcode.Decode(fi)
			if err != nil {
				panic(err)
			}
			fmt.Println(qrmatrix.Content)
		},
	}

	rootCmd.AddCommand(cmd)
}
