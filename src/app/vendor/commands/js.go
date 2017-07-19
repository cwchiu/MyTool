package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
	"github.com/spf13/cobra"
)

// https://github.com/robertkrimen/otto/blob/master/otto/main.go
func SetupJsCommand(rootCmd *cobra.Command) {
	var use_underscore bool
	cmd := &cobra.Command{
		Use:   "js [filename]",
		Short: "執行Javascript",
		Run: func(cmd *cobra.Command, args []string) {
			if !use_underscore {
				underscore.Disable()
			}

			err := func() error {
				var src []byte
				var err error
				if len(args) > 0 {
					src, err = ioutil.ReadFile(args[0])
				} else {
					src, err = ioutil.ReadAll(os.Stdin)
				}
				if err != nil {
					return err
				}

				vm := otto.New()
				_, err = vm.Run(src)
				return err
			}()
			if err != nil {
				switch err := err.(type) {
				case *otto.Error:
					fmt.Print(err.String())
				default:
					panic(err)
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&use_underscore, "underscore", "u", false, "啟用 underscore?")
	rootCmd.AddCommand(cmd)
}
