package date

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func SetupS2TCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "s2t <date string>",
		Short: "date string to timestamp",
		Long:  `date string to timestamp, format: 2006-01-02T15:04:05-0700`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                panic("need <date string>")
            }
            date_fmt := "2006-01-02T15:04:05-0700"
            dt, err := time.Parse(date_fmt, args[0])
            if err != nil {
                panic(err)
            }
            
			 // ISO8601
            fmt.Printf("Local: %v\n", dt.Local().Format(date_fmt))
            fmt.Printf("UTC: %v\n", dt.UTC().Format(date_fmt))
            fmt.Printf("Timestamp: %v\n", dt.UTC().Unix())
		},
	}

	rootCmd.AddCommand(cmd)
}
