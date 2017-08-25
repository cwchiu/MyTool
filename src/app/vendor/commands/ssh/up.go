package ssh

import (
    "log"
	"github.com/spf13/cobra"
    
)



func SetupUploadCommand(rootCmd *cobra.Command) {
	var username string
	var password string
	var host string
	var port int
	cmd := &cobra.Command{
		Use:   "up <remote-target> <local-src>",
		Short: "上傳檔案",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 2 {
                panic("required <remote-dir> <local-file>")
            }
            client, err := newftpclient(username, password, host, port)
            if err != nil {
                panic(err)
            }
            
            
            // 
            for _, local_fn := range(args[1:]) {
                log.Printf(">> Upload %s", local_fn)
                err = upload(client, local_fn, args[0])
                if err == nil {
                    log.Print("Ok")
                }else{
                    log.Print(err)
                }
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
