package ssh

import (
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

func SetupTtyCommand(rootCmd *cobra.Command) {
	var username string
	var password string
	var host string
	var port int
	cmd := &cobra.Command{
		Use:   "tty",
		Short: "terminal",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := connect(username, password, host, port)
			if err != nil {
				panic(err)
			}

			session, err := client.NewSession()
			if err != nil {
				panic(err)
			}
			defer session.Close()
			fd := int(os.Stdin.Fd())
			oldState, err := terminal.MakeRaw(fd)
			if err != nil {
				panic(err)
			}
			defer terminal.Restore(fd, oldState)

			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			session.Stdin = os.Stdin

			termWidth, termHeight, err := terminal.GetSize(fd)
			if err != nil {
				panic(err)
			}
			// Set up terminal modes
			modes := ssh.TerminalModes{
				ssh.ECHO:          1,     // enable echoing
				ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
				ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
			}

			// Request pseudo terminal
			if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
				log.Fatal(err)
			}

			session.Run("top")
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&username, "username", "u", "", "帳號")
	flags.StringVarP(&password, "password", "k", "", "密碼")
	flags.StringVarP(&host, "host", "t", "", "目的主機ip")
	flags.IntVarP(&port, "port", "p", 22, "目的主機port")
	rootCmd.AddCommand(cmd)
}
