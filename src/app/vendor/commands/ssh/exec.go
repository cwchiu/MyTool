package ssh

import (
	"github.com/spf13/cobra"
    "os"
)

func SetupExecCommand(rootCmd *cobra.Command) {
	var username string
	var password string
	var host string
	var port int
	cmd := &cobra.Command{
		Use:   "exec <cmd>",
		Short: "遠端執行命令",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                panic("required <cmd>")
            }
            client, err := connect(username, password, host, port)
            if err != nil {
                panic(err)
            }
            
            session, err := client.NewSession()
            if err != nil {
                panic(err)
            }
            defer session.Close()
            
            session.Stdout = os.Stdout
            session.Stderr = os.Stderr
            
            for _, input := range(args) {
                session.Run(input)
            }
            
		},
	}
    flags := cmd.Flags()
	flags.StringVarP(&username, "username", "u", "", "帳號")
	flags.StringVarP(&password, "password", "k", "", "密碼")
	flags.StringVarP(&host, "host", "t", "", "目的主機ip")
	flags.IntVarP(&port, "port", "p", 22, "目的主機port")
	rootCmd.AddCommand(cmd)
}
