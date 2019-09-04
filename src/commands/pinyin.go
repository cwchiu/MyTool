package commands

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/spf13/cobra"
)

func init() {
	var mode int
	cmd := &cobra.Command{
		Use:   "pinyin <word>",
		Short: "中文轉拼音",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <word>")
			}
			a := pinyin.NewArgs()
			if mode == 1 {
				a.Style = pinyin.Tone
			} else if mode == 2 {
				a.Style = pinyin.Tone2
			}
			ret := pinyin.Pinyin(args[0], a)
			if len(ret) > 0 {
				for _, word := range ret {
					fmt.Printf("%v ", word[0])
				}
				fmt.Println()
			}
		},
	}
	cmd.Flags().IntVarP(&mode, "mode", "m", 0, "0=預設, 1=包含聲調, 2=聲調用數字表示")
	rootCmd.AddCommand(cmd)
}
