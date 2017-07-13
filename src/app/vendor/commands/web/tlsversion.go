package web

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"regexp"
)

func fetchTLSVersion(url string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	switch resp.TLS.Version {
	case tls.VersionSSL30:
		return "SSLV3", nil
	case tls.VersionTLS10:
		return "TLSv1", nil
	case tls.VersionTLS11:
		return "TLSv1.1", nil
	case tls.VersionTLS12:
		return "TLSv1.2", nil
	default:
		return "Unknown", nil
	}
}
func SetupTlsVersionCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "tls <url>",
		Short: "tls version",
		Long:  `tls version`,
		Run: func(cmd *cobra.Command, args []string) {
			for _, url := range args {
				// url := "httPs://www.baidu.com"
				matched, err := regexp.MatchString("(?i)^https", url)
				if matched == false || err != nil {
					fmt.Printf("%s skip\n", url)
					continue
				}
				ver, err := fetchTLSVersion(url)
				if err != nil {
					fmt.Printf("%s %s\n", url, err)
				} else {
					fmt.Printf("%s %s\n", url, ver)
				}

			}
		},
	}

	rootCmd.AddCommand(cmd)

}
