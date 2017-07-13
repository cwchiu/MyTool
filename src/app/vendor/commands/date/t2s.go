package date

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
    "regexp"
    "strconv"
)

func SetupT2SCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "t2s <timestamp>",
		Short: "timestamp to date string",
		Long:  `timestamp to date string`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                panic("need timestamp")
            }
            
            matched, regex_err := regexp.MatchString(`^\d{1,10}$`, args[0])
            if matched == false || regex_err != nil {
                panic("input is not valid timestamp")
            }
            
            ts, err := strconv.ParseInt(args[0], 10, 64)
            if err != nil {
                panic("input is not valid timestamp")
            }
            
            t := time.Unix(ts, 0)
            
			date_fmt := "2006-01-02T15:04:05-0700" // ISO8601
            fmt.Printf("Local: %v\n", t.Local().Format(date_fmt))
            fmt.Printf("UTC: %v\n", t.UTC().Format(date_fmt))
            fmt.Printf("Timestamp: %v\n", t.UTC().Unix())
		},
	}

	rootCmd.AddCommand(cmd)
}
