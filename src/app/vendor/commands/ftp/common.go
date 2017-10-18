package ftp

import (
	ftpclient "github.com/jlaffaye/ftp"
	"net/url"
	// "fmt"
)

type FtpInfo struct {
	Addr     string
	Username string
	Password string
	Path     string
}

func ParseUrl(v string) *FtpInfo {
	urlinfo, err := url.Parse(v)
	msg := "invalid <remote file> format, ex: ftp://127.0.0.1/myfolder/file.txt"
	if err != nil {
		panic(msg)
	}

	addr := urlinfo.Hostname()
	if len(addr) == 0 {
		panic(msg)
	}
	if urlinfo.Port() != "" {
		addr += ":" + urlinfo.Port()
	}

	username := "anonymous"
	password := "anonymous"

	if urlinfo.User != nil {
		username = urlinfo.User.Username()
		v, b := urlinfo.User.Password()
		if b {
			password = v
		}
	}
	// fmt.Println(addr)
	// fmt.Println(username)
	// fmt.Println(password)
	// fmt.Println(urlinfo.Path)
	return &FtpInfo{
		Username: username,
		Password: password,
		Addr:     addr,
		Path:     urlinfo.Path,
	}
}

func CreateFtpClient(info *FtpInfo) *ftpclient.ServerConn {
	conn, err := ftpclient.Connect(info.Addr)
	if err != nil {
		panic(err)
	}

	err = conn.Login(info.Username, info.Password)
	if err != nil {
		panic(err)
	}
	return conn
}
