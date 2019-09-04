package server

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
	"io"
	"os"
	"unsafe"
	"os/exec"
	"syscall"
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func SSHSessionHandle(s ssh.Session) {
	ptyReq, winCh, isPty := s.Pty()
	if isPty {
		cmd := exec.Command("top")
		cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
		f, err := pty.Start(cmd)
		if err != nil {
			panic(err)
		}
		go func() {
			for win := range winCh {
				setWinsize(f, win.Width, win.Height)
			}
		}()
		go func() {
			io.Copy(f, s) // stdin
		}()
		io.Copy(s, f) // stdout
	} else {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	}
}
