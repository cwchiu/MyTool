package date

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func SetupNowCommand(rootCmd *cobra.Command) {
	var ts_only bool
	cmd := &cobra.Command{
		Use:   "now",
		Short: "now datetime",
		Long:  `now datetime`,
		Run: func(cmd *cobra.Command, args []string) {
			now := time.Now()
			date_fmt := "2006-01-02T15:04:05-0700" // ISO8601
			if ts_only {
				fmt.Println(now.UTC().Unix())
			} else {
				fmt.Printf("Local: %v\n", now.Local().Format(date_fmt))
				fmt.Printf("UTC: %v\n", now.UTC().Format(date_fmt))
				fmt.Printf("Timestamp: %v\n", now.UTC().Unix())
			}
		},
	}

	cmd.Flags().BoolVarP(&ts_only, "ts", "t", false, "timestamp only")
	rootCmd.AddCommand(cmd)
}
