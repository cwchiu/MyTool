package server

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/spf13/cobra"
	"io"
	"log"
)

func SetupSSHCommand(rootCmd *cobra.Command) {
	var port int32
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "ssh Server",
		Long:  `ssh Server`,
		Run: func(cmd *cobra.Command, args []string) {

			ssh.Handle(func(s ssh.Session) {
				io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
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
