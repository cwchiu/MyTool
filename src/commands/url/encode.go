package url

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
)

func thunder(url string) string {
	return "thunder://" + base64.StdEncoding.EncodeToString([]byte("AA"+url+"ZZ"))
}

func flashget(url string) string {
	return "flashget://" + base64.StdEncoding.EncodeToString([]byte("[FLASHGET]"+url+"[FLASHGET]")) + "&aiyh"
}
func qq(url string) string {
	return "qqdl://" + base64.StdEncoding.EncodeToString([]byte(url))
}

func SetupEncodeCommand(rootCmd *cobra.Command) {
	var encode string
	var list bool
	cmd := &cobra.Command{
		Use:   "encode <url>",
		Short: "產生指定的下載編碼格式",
		Run: func(cmd *cobra.Command, args []string) {
			if list {
				fmt.Println("encode: flashget, thunder, qq")
				return
			}

			if len(args) < 1 {
				panic("required <url>")
			}
			switch encode {
			case "qq":
				fmt.Println(qq(args[0]))
			case "thunder":
				fmt.Println(thunder(args[0]))
			default:
				fmt.Println(flashget(args[0]))
			}
		},
	}
	cmd.Flags().StringVarP(&encode, "encode", "e", "flashget", "編碼格式")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "列出支援的編碼格式")
	rootCmd.AddCommand(cmd)
}
