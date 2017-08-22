package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Exists(fn string) bool {
	if _, err := os.Stat(fn); err == nil {
		return true
	}
	return false
}

// openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
func SetupStaticCommand(rootCmd *cobra.Command) {
	var port int32
	var root string
	var filename_cert string
	var filename_key string
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

			upload_folder := filepath.Join(root, "upload")
			os.Mkdir(upload_folder, 0777)

			http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
				log.Print("==== New Request ====")
				log.Printf("Http Method: %s", r.Method)
				log.Printf("Protocol: %s", r.Proto)
				log.Printf("Host: %s", r.Host)
				log.Printf("RequestURI: %s", r.RequestURI)
				log.Printf("RemoteAddress: %s", r.RemoteAddr)
				// log.Printf("Cookies: %v", r.Cookies())
				log.Print("Headers: ")
				for k, v := range r.Header {
					log.Printf("  %s: %s", k, v)
				}
				if r.Method == "POST" {
					log.Printf("ContentLength: %d", r.ContentLength)
					body, err := ioutil.ReadAll(r.Body)
					if err == nil {
						log.Printf("Body: %s", string(body))
					} else {
						log.Print(err)
					}
				}
			})

			http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				//POST takes the uploaded file(s) and saves it to disk.
				case "POST":
					//parse the multipart form in the request
					err := r.ParseMultipartForm(100000)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					//get a ref to the parsed multipart form
					m := r.MultipartForm

					//get the *fileheaders
					files := m.File["uploadfile"]
					for i, _ := range files {
						//for each fileheader, get a handle to the actual file
						file, err := files[i].Open()
						defer file.Close()
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						//create destination file making sure the path is writeable.
						target_fn := filepath.Join(upload_folder, files[i].Filename)
						fmt.Println(filepath.Base(target_fn))
						dst, err := os.Create(target_fn)
						defer dst.Close()
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						//copy the uploaded file to the destination file
						if _, err := io.Copy(dst, file); err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

					}

				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			})
			http.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir(root))))
			if Exists(filename_cert) && Exists(filename_key) {
				protocol = "https"
			}
			log.Printf("Listen Port: %d\n", port)
			log.Printf("Home Folder: %s\n", root)
			log.Printf("Browser: %s://127.0.0.1:%d/s/", protocol, port)
			log.Printf("Upload: %s://127.0.0.1:%d/upload", protocol, port)
			log.Printf("API: %s://127.0.0.1:%d/api", protocol, port)

			if protocol == "https" {
				err = http.ListenAndServeTLS(fmt.Sprintf(":%d", port), filename_cert, filename_key, nil)
			} else {
				err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
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
	rootCmd.AddCommand(cmd)
}
