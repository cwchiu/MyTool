package clamav

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

type ClamAV struct {
	Entry string
}

func (self ClamAV) request(command string) (string, error) {
	conn, err := net.Dial("tcp", self.Entry)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	conn.Write([]byte(command))
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (self ClamAV) Version() (string, error) {
	return self.request("VERSION")
}

func (self ClamAV) Scan(fn string) (string, error) {
	ret, err := self.request(fmt.Sprintf("SCAN %s", fn))
	if err != nil {
		return "", err
	}

	if strings.Contains(ret, "failed: No such file or directory") {
		return "", errors.New("No such file or directory")
	}
	return strings.TrimSpace(strings.SplitN(ret, ":", 2)[1]), nil
}

func CreateClamAV(entry string) (*ClamAV, error) {
	return &ClamAV{
		Entry: entry,
	}, nil
}
