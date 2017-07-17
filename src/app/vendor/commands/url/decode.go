package url

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	// "io"
	"fmt"
	"strings"
)

func decode(url string) string {
	urlodd := strings.Split(url, "//")
	head := strings.ToLower(urlodd[0])
	behind := urlodd[1]
	switch head {
	case "thunder:":
		data, err := base64.StdEncoding.DecodeString(behind)
		if err != nil {
			panic(err)
		}
		return string(data[2 : len(data)-2])
		// $url=substr(base64_decode($behind), 2, -2);//base64解密，去掉前面的AA和後面ZZ
	case "flashget:":
		url1 := strings.SplitN(behind, "&", 2)
		// $url1=explode('&',$behind,2);
		data, err := base64.StdEncoding.DecodeString(url1[0])
		if err != nil {
			panic(err)
		}
		return string(data[10 : len(data)-10])
		// $url=substr(base64_decode($url1[0]), 10, -10);//base64解密，去掉前面後的[FLASHGET]
	case "qqdl:":
		// $url=base64_decode($behind);//base64解密
		data, err := base64.StdEncoding.DecodeString(behind)
		if err != nil {
			panic(err)
		}
		return string(data)
	default:
		return url
	}
}
func SetupDecodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "URL解碼",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <url>")
			}

			fmt.Println(decode(args[0]))
		},
	}

	rootCmd.AddCommand(cmd)
}
