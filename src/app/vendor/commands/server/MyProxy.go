package server

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

func SetupProxyCommand(rootCmd *cobra.Command) {
	var port int32
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Proxy Server",
		Long:  `proxy Server`,
		Run: func(cmd *cobra.Command, args []string) {

			proxy := goproxy.NewProxyHttpServer()
			proxy.Verbose = true
			log.Printf("Listen Port: %d\n", port)

			err := http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
			if err != nil {
				panic(err)
			}

		},
	}
	cmd.Flags().Int32VarP(&port, "port", "p", 18080, "listen port")
	rootCmd.AddCommand(cmd)
}
