package web

import (
	"fmt"
	"github.com/spf13/cobra"
    "net/http"
    "net/url"
    "io"
    "strings"
    "os"
)

func SetupDownloadCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "download <url>",
		Short: "download file",
		Long:  `download file`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                panic("required <url>")
            }
            
            urlinfo, err := url.Parse(args[0])
            if err != nil {
                panic(err)
            }
            tokens := strings.Split(urlinfo.Path, "/")
            var filename string
            if len(tokens) == 0 {
                filename = "unknown.html"
            }else{
                filename = tokens[len(tokens)-1]
            }
            output, err := os.Create(filename)
            if err != nil {
                panic(err)
            }
            defer output.Close()
            
            req, err := http.NewRequest("GET", args[0], nil)
            if err != nil {
                panic(err)
            }
            
            req.Header.Add("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")
            
            client := &http.Client{}
            resp, err := client.Do(req)
            if err != nil {
                panic(err)
            }
            
            defer resp.Body.Close()
            
            n, download_err := io.Copy(output, resp.Body)
            if download_err != nil {
                panic(download_err)
            }
            fmt.Printf("Filename: %s\n", filename)
            fmt.Printf("Size: %d\n", n)
		},
	}

	rootCmd.AddCommand(cmd)

}
