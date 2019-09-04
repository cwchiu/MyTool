package web

import (
	"fmt"
	pastebin "github.com/cwchiu/go-pastebin"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func SetupPasteBinCommand(rootCmd *cobra.Command) {

	cmd := &cobra.Command{
		Use:   "pastebin <title>",
		Short: "新增 pastebin",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <title>")
			}
			p := pastebin.Pastebin{Key: "7ab028078cd3dbe433ded35eeb8786e4"}

			code, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}

			pid, err := p.Put(*pastebin.CreateNormalPaste(args[0], string(code)))
			if err != nil {
				panic(err)
			}
			fmt.Println(pid)
		},
	}

	rootCmd.AddCommand(cmd)

}
