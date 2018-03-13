package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"libs/opencc"
)

func setupOpenCCS2TCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "s2t <Simplified Text>",
		Short: "簡轉繁",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need <Simplified Text>")
			}
            
            s, err := opencc.ToTraditional(args[0])
            
            if err != nil {
                panic(err)
            }
            
            fmt.Println(s)
		},
	}

	rootCmd.AddCommand(cmd)
}

func setupOpenCCT2SCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "t2s <Traditional Text>",
		Short: "繁轉簡",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need <Traditional Text>")
			}
            
            s, err := opencc.ToSimplified(args[0])
            
            if err != nil {
                panic(err)
            }
            
            fmt.Println(s)
		},
	}

	rootCmd.AddCommand(cmd)
}


func init() {
	cmd := &cobra.Command{Use: "opencc", Short: "繁/簡轉換"}

	setupOpenCCS2TCommand(cmd)
	setupOpenCCT2SCommand(cmd)

	rootCmd.AddCommand(cmd)
}
