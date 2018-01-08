package server

import (
	"errors"
	"github.com/bradfitz/go-smtpd/smtpd"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type env struct {
	*smtpd.BasicEnvelope
}

func (e *env) AddRecipient(rcpt smtpd.MailAddress) error {
	if strings.HasPrefix(rcpt.Email(), "bad@") {
		return errors.New("we don't send email to bad@")
	}
	return e.BasicEnvelope.AddRecipient(rcpt)
}

func onNewMail(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	log.Printf("ajas: new mail from %q", from)
	return &env{new(smtpd.BasicEnvelope)}, nil
}
func onNewConnection(c smtpd.Connection) error {
	log.Println(c)
	return nil
}

func SetupSMTPCommand(rootCmd *cobra.Command) {
	var addr string
	cmd := &cobra.Command{
		Use:   "smtp",
		Short: "Fake SMTP Server",
		Run: func(cmd *cobra.Command, args []string) {

			s := &smtpd.Server{
				Addr:            addr,
				PlainAuth:       true,
				OnNewMail:       onNewMail,
				OnNewConnection: onNewConnection,
			}
            log.Printf("Listen: %s", addr)
			err := s.ListenAndServe()
			if err != nil {
				log.Fatalf("ListenAndServe: %v", err)
			} 

		},
	}

	cmd.Flags().StringVarP(&addr, "addr", "a", ":25", "Listen Port")
	rootCmd.AddCommand(cmd)
}
