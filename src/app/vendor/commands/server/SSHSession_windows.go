package server

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
)

func SSHSessionHandle(s ssh.Session) {
	io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
}
