package fs

import (
	// "fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func SetupCpCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "cp",
		Short: "copy file from srouce to target",
		Long:  `copy file from srouce to target`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("need an filename")
			}
			source, err := ioutil.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			for _, fn := range args[1:] {
				err = ioutil.WriteFile(fn, source, 0777)
				if err != nil {
					panic(err)
				}
			}

		},
	}

	rootCmd.AddCommand(cmd)
}
