package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	m["Method"] = r.Method
	m["Protocol"] = r.Proto
	m["RemoteAddr"] = r.RemoteAddr
	m["RequestURI"] = r.RequestURI
	m["Header"] = r.Header

	if r.Method == "POST" {
		m["ContentLength"] = r.ContentLength

		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			m["Body"] = string(body)
		} else {
			log.Print(err)
		}
	}

	bs, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(bs)
}

var user_response map[string]interface{}

func init() {
	user_response = make(map[string]interface{})
	user_response["body"] = "Hello"
	user_response["code"] = 200
	user_response["mime"] = "text/plain; charset=utf-8"
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var form map[string]string
	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	m := make(map[string]interface{})
	mime_type, exists := form["MimeType"]
	if exists == false || m["mime"] == "" {
		mime_type = "text/plain"
	}
	m["mime"] = mime_type

	code, exists := form["code"]
	if exists == false {
		code = "200"
	}
	code_i, err := strconv.Atoi(code)
	if err != nil {
		code_i = 200
	}
	m["code"] = code_i
	m["body"] = form["body"]
	user_response = m
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", user_response["mime"].(string))
	w.WriteHeader(user_response["code"].(int))
	w.Write([]byte(user_response["body"].(string)))
}

func HttpCodeResponseHandle(w http.ResponseWriter, r *http.Request) {
	urlinfo, err := url.Parse(r.URL.String())
	code_i := 200
	if err == nil {
		q := urlinfo.Query()
		code := q.Get("v")
		code_i, err = strconv.Atoi(code)
		if err != nil {
			code_i = 200
		}
	}

	w.WriteHeader(code_i)
}

func RegisterApi() {
	http.HandleFunc("/api/register", AddHandler)
	http.HandleFunc("/api/test", TestHandler)
	http.HandleFunc("/api/request", ApiHandler)
	http.HandleFunc("/api/code", HttpCodeResponseHandle)
}
