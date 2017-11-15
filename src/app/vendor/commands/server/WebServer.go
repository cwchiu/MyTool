package server

import (
	// "encoding/base64"
	// "bytes"
	"encoding/json"
	"fmt"
	"github.com/abbot/go-http-auth"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
    "regexp"
	"time"
	// "strings"
)

func exists(fn string) bool {
	if _, err := os.Stat(fn); err == nil {
		return true
	}
	return false
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
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
}

func createUploadHandler(root string) http.HandlerFunc {
	upload_folder := filepath.Join(root, "upload")
	os.Mkdir(upload_folder, 0777)
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func pass(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
}

func createDigestAuthHandler() http.HandlerFunc {
	digest_auth := auth.NewDigestAuthenticator("chuiwenchiu.wordpress.com", func(user, realm string) string {
		if user == "guest" {
			return "1234"
		}
		return ""
	})
	digest_auth.PlainTextSecrets = true
	return digest_auth.Wrap(pass)
}

func createBasicAuthHandler() http.HandlerFunc {
	basic_auth := auth.NewBasicAuthenticator("chuiwenchiu.wordpress.com", func(user, realm string) string {
		log.Printf("%v", realm)
		if user == "guest" {
			// hello
			password := "1234"
			magic := "$1$" // 前後一定要有 $
			salt := "dlPL2MqE"

			// hashedPassword := []byte("$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1")
			// parts := bytes.SplitN(hashedPassword, []byte("$"), 4)
			// magic2 := []byte("$" + string(parts[1]) + "$")
			// salt2 := parts[2]
			// fmt.Printf("%v = %v\n", string(magic2), string(magic))
			// fmt.Printf("%v = %v\n", string(salt2), string(salt))
			// return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
			// v := magic + salt  + "$" + string(auth.MD5Crypt([]byte(password), []byte(salt), []byte(magic)))
			// fmt.Printf(v)
			return string(auth.MD5Crypt([]byte(password), []byte(salt), []byte(magic)))
		}
		return ""
	})

	return basic_auth.Wrap(pass)
}

// https://gist.github.com/thealexcons/4ecc09d50e6b9b3ff4e2408e910beb22
type userCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtClaims struct {
	Username string `json:"username"`
	// recommended having
	jwt.StandardClaims
}

type tokenData struct {
	Token string `json:"token"`
}

const JWT_SECRET string = "!@#$%^&"

func jsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// curl -k -H "Content-Type: application/json" -X POST -d '{"username":"guest","password":"1234"}' https://127.0.0.1:28080/auth/jwt/login
// curl -k -H "Content-Type: application/json" -X POST -d '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imd1ZXN0IiwiZXhwIjoxNTA1NzEwNDUxLCJpc3MiOiJsb2NhbGhvc3Q6OTAwMCJ9.Pgaz0u3XkDlSqfsAJeCzRVJqsfmYS89wIeUIKmebyNY"}' https://127.0.0.1:28080/auth/jwt/test
func jwtLogin(w http.ResponseWriter, r *http.Request) {

	var user userCredentials

	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	log.Println(user.Username, user.Password)

	//validate user credentials
	if user.Username != "guest" || user.Password != "1234" {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	// Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	// expireCookie := time.Now().Add(time.Hour * 1)

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := jwtClaims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte(JWT_SECRET))
	// cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	// http.SetCookie(w, &cookie)

	//create a token instance using the token string
	response := tokenData{signedToken}
	jsonResponse(response, w)
}

func jwtTest(w http.ResponseWriter, r *http.Request) {
	var data tokenData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	// Return a Token using the cookie
	token, err := jwt.ParseWithClaims(data.Token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		fmt.Println(claims)
		fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", claims.Username)

	} else {
		http.NotFound(w, r)
		return
	}
}


// openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
func SetupWebCommand(rootCmd *cobra.Command) {
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

			http.HandleFunc("/api", apiHandler)
			http.HandleFunc("/upload", createUploadHandler(root))
			http.HandleFunc("/auth/jwt/login", jwtLogin)
			http.HandleFunc("/auth/jwt/test", jwtTest)
			http.HandleFunc("/auth/basic", createBasicAuthHandler())
			http.HandleFunc("/auth/digest", createDigestAuthHandler())

            fileserver := http.StripPrefix("/s/", http.FileServer(http.Dir(root)))
            
            wasm := regexp.MustCompile("\\.wasm$")
			http.HandleFunc("/s/", func (w http.ResponseWriter, r *http.Request) {
                ruri := r.RequestURI
                if wasm.MatchString(ruri) {
                    w.Header().Set("Content-Type", "application/wasm")
                }
                fileserver.ServeHTTP(w, r)
            })
			if exists(filename_cert) && exists(filename_key) {
				protocol = "https"
			}
			log.Printf("Listen Port: %d\n", port)
			log.Printf("Home Folder: %s\n", root)
			log.Printf("Browser: %s://127.0.0.1:%d/s/", protocol, port)
			log.Printf("Upload: %s://127.0.0.1:%d/upload", protocol, port)
			log.Printf("Basic Auth(guest/1234): %s://127.0.0.1:%d/auth/basic", protocol, port)
			log.Printf("Digest Auth(guest/1234): %s://127.0.0.1:%d/auth/digest", protocol, port)
			log.Printf("JWT Auth(guest/1234): %s://127.0.0.1:%d/auth/jwt/login", protocol, port)
			log.Printf("JWT Test: %s://127.0.0.1:%d/auth/jwt/test", protocol, port)
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
