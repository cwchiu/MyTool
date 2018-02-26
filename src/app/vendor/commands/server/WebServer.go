package server

import (
	"commands/server/web"
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func exists(fn string) bool {
	if _, err := os.Stat(fn); err == nil {
		return true
	}
	return false
}

// openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
func SetupWebCommand(rootCmd *cobra.Command) {
	var port int32
	var root string
	var filename_cert string
	var filename_key string
	var open_browser bool
	protocol := "http"
	cmd := &cobra.Command{
		Use:   "web",
		Short: "Http Web Server",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := os.Stat(root)
			if err != nil {
				root = "."
			}

			root, err = filepath.Abs(root)
			if err != nil {
				root = "."
			}

			web.RegisterApi()
			web.RegisterFile(root)
			web.RegisterAuth()
			web.RegisterCaptacha()
			web.RegisterMimeHandler()

			fileserver := http.StripPrefix("/s/", http.FileServer(http.Dir(root)))

			wasm := regexp.MustCompile("\\.wasm$")
			http.HandleFunc("/s/", func(w http.ResponseWriter, r *http.Request) {
				ruri := r.RequestURI
				if wasm.MatchString(ruri) {
					w.Header().Set("Content-Type", "application/wasm")
				}
				fileserver.ServeHTTP(w, r)
			})
			if exists(filename_cert) && exists(filename_key) {
				protocol = "https"
			}
			api_gateway := fmt.Sprintf("%s://127.0.0.1:%d", protocol, port)
			log.Printf("Listen Port: %d\n", port)
			log.Printf("Home Folder: %s\n", root)
			log.Printf("Browser: %s/s/", api_gateway)
			log.Printf("Upload: %s/upload", api_gateway)
			log.Printf("Basic Auth(guest/1234): %s/auth/basic", api_gateway)
			log.Printf("Digest Auth(guest/1234): %s/auth/digest", api_gateway)
			log.Printf("JWT Auth(guest/1234): %s/auth/jwt/login", api_gateway)
			log.Printf("JWT Test: %s/auth/jwt/test", api_gateway)
			log.Printf("captcha Get: %s/captcha/verify", api_gateway)
			log.Printf("captcha verify: %s/captcha/get", api_gateway)
			log.Printf("API: %s/api", api_gateway)

			if open_browser {
				c1 := make(chan string, 1)
				select {
				case res := <-c1:
					fmt.Println(res)
				case <-time.After(time.Second * 1):
					fmt.Println("timeout 1")
					browser.OpenURL(fmt.Sprintf("%s/auth/basic", api_gateway))
				}
			}
			server := &http.Server{
				Addr:           fmt.Sprintf(":%d", port),
				ReadTimeout:    30 * time.Second,
				WriteTimeout:   30 * time.Second,
				MaxHeaderBytes: 1 << 20}
			if protocol == "https" {
				err = server.ListenAndServeTLS(filename_cert, filename_key)
			} else {
				err = server.ListenAndServe()
			}

			if err != nil {
				panic(err)
			}

		},
	}
	cmd.Flags().Int32VarP(&port, "port", "p", 28080, "listen port")
	cmd.Flags().StringVarP(&root, "root", "r", ".", "root folder")
	cmd.Flags().StringVarP(&filename_cert, "tls-cert", "c", "cert.pem", "cert.pem")
	cmd.Flags().StringVarP(&filename_key, "tls-key", "k", "key.pem", "key.pem")
	cmd.Flags().BoolVarP(&open_browser, "browser", "b", false, "瀏覽器自動開啟")
	rootCmd.AddCommand(cmd)
}
