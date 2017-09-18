package date

import (
	"fmt"
	"github.com/spf13/cobra"
    "github.com/nosixtools/solarlunar" 
    "time"
)

func SetupS2LCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "s2l <date string>",
		Short: "陽曆轉陰曆, ex: 2006-01-02",
		Run: func(cmd *cobra.Command, args []string) {
            now := time.Now()
			date_fmt := "2006-01-02" // ISO8601
            solarDate := now.Local().Format(date_fmt)
            if len(args) > 0 {
                solarDate = args[0]
            }
            fmt.Println(solarlunar.SolarToChineseLuanr(solarDate))
            fmt.Println(solarlunar.SolarToSimpleLuanr(solarDate))
            
		},
	}

	rootCmd.AddCommand(cmd)
}
