package server

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os/exec"
	"syscall"
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func SetupSSHCommand(rootCmd *cobra.Command) {
	var port int32
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "ssh Server",
		Long:  `ssh Server`,
		Run: func(cmd *cobra.Command, args []string) {

			ssh.Handle(func(s ssh.Session) {
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
			})
			log.Printf("Listen %d\n", port)
			err := ssh.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
			if err != nil {
				panic(err)
			}

		},
	}
	cmd.Flags().Int32VarP(&port, "port", "p", 22, "listen port")
	rootCmd.AddCommand(cmd)
}
