package web

import (
	"encoding/json"
	"fmt"
	"github.com/abbot/go-http-auth"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

func pass(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
}

func CreateDigestAuthHandler() http.HandlerFunc {
	digest_auth := auth.NewDigestAuthenticator("chuiwenchiu.wordpress.com", func(user, realm string) string {
		if user == "guest" {
			return "1234"
		}
		return ""
	})
	digest_auth.PlainTextSecrets = true
	return digest_auth.Wrap(pass)
}

func CreateBasicAuthHandler() http.HandlerFunc {
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
func JWTLogin(w http.ResponseWriter, r *http.Request) {

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

func JWTTest(w http.ResponseWriter, r *http.Request) {
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

func RegisterAuth() {
	http.HandleFunc("/auth/jwt/login", JWTLogin)
	http.HandleFunc("/auth/jwt/test", JWTTest)
	http.HandleFunc("/auth/basic", CreateBasicAuthHandler())
	http.HandleFunc("/auth/digest", CreateDigestAuthHandler())
}
