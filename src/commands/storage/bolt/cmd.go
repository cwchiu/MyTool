package bolt

import (
	"fmt"
	"github.com/spf13/cobra"
	lib "github.com/cwchiu/MyTool/libs/bolt"
)

func setupPutCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "put <database> <bucket> <key> <value>",
		Short: "儲存",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 4 {
				panic("<database> <bucket> <key> <value>")
			}
			err := lib.Put(args[0], args[1], args[2], args[3])
			if err != nil {
				panic(err)
			}

			fmt.Println("Ok")
		},
	}
	rootCmd.AddCommand(cmd)
}

func setupGetCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "get <database> <bucket> <key>",
		Short: "取值",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				panic("<database> <bucket> <key>")
			}
			value, err := lib.Get(args[0], args[1], args[2])
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(value)
		},
	}
	rootCmd.AddCommand(cmd)
}

func setupScanCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "scan <database> <bucket>",
		Short: "掃描 bucket",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("<database> <bucket>")
			}
			err := lib.Scan(args[0], args[1], func(_, key, value string) error {
				fmt.Printf("%s:%s\n", key, value)

				return nil
			})

			if err != nil {
				panic(err)
			}

		},
	}
	rootCmd.AddCommand(cmd)
}

func setupListBucketCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "buckets <database>",
		Short: "bucket 列表",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("<database>")
			}
			err := lib.ListBucket(args[0], func(name string) error {
				fmt.Println(name)

				return nil
			})

			if err != nil {
				panic(err)
			}

		},
	}
	rootCmd.AddCommand(cmd)
}

func SetupCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{Use: "bolt", Short: "bolt"}

	setupPutCommand(cmd)
	setupGetCommand(cmd)
	setupScanCommand(cmd)
	setupListBucketCommand(cmd)

	rootCmd.AddCommand(cmd)
}
