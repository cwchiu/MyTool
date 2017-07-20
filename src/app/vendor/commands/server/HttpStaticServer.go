package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func SetupStaticCommand(rootCmd *cobra.Command) {
	var port int32
	var root string
	cmd := &cobra.Command{
		Use:   "static",
		Short: "Static Http Server",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := os.Stat(root)
			if err != nil {
				root = "."
			}

			root, err = filepath.Abs(root)
			if err != nil {
				root = "."
			}

			log.Printf("Listen Port: %d\n", port)
			log.Printf("Home Folder: %s\n", root)

			err = http.ListenAndServe(fmt.Sprintf(":%d", port),
				http.FileServer(http.Dir(root)))
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().Int32VarP(&port, "port", "p", 28080, "listen port")
	cmd.Flags().StringVarP(&root, "root", "r", ".", "root folder")
	rootCmd.AddCommand(cmd)
}
