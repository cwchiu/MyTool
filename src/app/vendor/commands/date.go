package commands

import (
	"fmt"
	"github.com/nosixtools/solarlunar"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"libs/date"
	"os"
	"regexp"
	"strconv"
	"time"
)

func printDate(dt *date.Date) {

	data := [][]string{
		[]string{"Local", dt.Local()},
		[]string{"ISO8601", dt.Iso8601()},
		[]string{"RFC850", dt.Rfc850()},
		[]string{"RFC1036", dt.Rfc1036()},
		[]string{"RFC1123", dt.Rfc1123()},
		[]string{"RFC2822", dt.Rfc2822()},
		[]string{"RFC3339", dt.Rfc3339()},
		[]string{"Ruby", dt.Ruby()},
		[]string{"ANSI C", dt.AnsiC()},
		[]string{"Unix", dt.Unix()},
		[]string{"Epoch", dt.Epoch()},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()

}

func setupS2LCommand(rootCmd *cobra.Command) {
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

func setupL2SCommand(rootCmd *cobra.Command) {
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

func setupT2SCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "t2s <timestamp>",
		Short: "timestamp to date string",
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

			dt, _ := date.FromUnix(ts)
			printDate(dt)
		},
	}

	rootCmd.AddCommand(cmd)
}

func setupS2TCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "s2t <date string>",
		Short: "date string to timestamp",
		Long:  `date string to timestamp, format: 2006-01-02T15:04:05-0700`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need <date string>")
			}
			dt, err := date.FromString(args[0])
			if err != nil {
				panic(err)
			}

			printDate(dt)
		},
	}

	rootCmd.AddCommand(cmd)
}

func setupNowCommand(rootCmd *cobra.Command) {
	var ts_only bool
	cmd := &cobra.Command{
		Use:   "now",
		Short: "now datetime",
		Long:  `now datetime`,
		Run: func(cmd *cobra.Command, args []string) {
			dt, _ := date.FromTime(time.Now())
			if ts_only {
				fmt.Println(dt.Epoch())
			} else {
				printDate(dt)
			}
		},
	}

	cmd.Flags().BoolVarP(&ts_only, "ts", "t", false, "timestamp only")
	rootCmd.AddCommand(cmd)
}

func init() {
	cmd := &cobra.Command{Use: "date", Short: "日期/時間"}

	setupT2SCommand(cmd)
	setupS2TCommand(cmd)

	setupS2LCommand(cmd)
	setupL2SCommand(cmd)

	setupNowCommand(cmd)

	rootCmd.AddCommand(cmd)
}
