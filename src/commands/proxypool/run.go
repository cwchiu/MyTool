package proxypool

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetupRunCommand(rootCmd *cobra.Command) {
	// var mgo_addr string
	// var mgo_db string
	// var mgo_col string
	var port int
	cmd := &cobra.Command{
		Use:   "run",
		Short: "啟動服務",
		Run: func(cmd *cobra.Command, args []string) {
			check()
			run(fmt.Sprintf(":%d", port))
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&port, "port", "p", 8080, "Listen Port")
	// flags.StringVarP(&mgo_addr, "mgo-addr", "a", "mongodb://127.0.0.1:27017?maxPoolSize=15", "MongoDB連線參數")
	// flags.StringVarP(&mgo_db, "mgo-db", "d", "proxypool", "MongoDB Database Name")
	// flags.StringVarP(&mgo_col, "mgo-col", "c", "pool", "MongoDB Collection name")

	rootCmd.AddCommand(cmd)
}
