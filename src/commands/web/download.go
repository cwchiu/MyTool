package web

import (
    "github.com/spf13/cobra"
    "github.com/cwchiu/MyTool/libs"
)

func SetupDownloadCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "download <url> [filename]",
        Long: `明確指定 filename 才能啟用續傳`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <url>")
			}
            
            filename := ""
            if len(args)>1 && args[1] != "" {
               filename = args[1]
            }
            
            err := libs.DownloadOne(args[0], filename)
            if err != nil {
                panic(err)
            }
		},
	}
	rootCmd.AddCommand(cmd)

}