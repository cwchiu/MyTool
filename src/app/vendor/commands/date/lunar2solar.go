package date

import (
	"fmt"
	"github.com/spf13/cobra"
    "github.com/nosixtools/solarlunar" 
)

func SetupL2SCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "l2s <date string>",
		Short: "陰曆轉陽曆, ex:2017-07-28",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                panic("need <date string>")
            }
            
            lunarDate := args[0]
            fmt.Println(solarlunar.LunarToSolar(lunarDate, false))
		},
	}

	rootCmd.AddCommand(cmd)
}
