package commands

import (
	"github.com/spf13/cobra"
    "fmt"
)

var	rootCmd = &cobra.Command{
    Use: "tool", 
    Long: fmt.Sprintf(`我的個人常用工具
Version: %s
Site: https://chuiwenchiu.wordpress.com
Github: https://github.com/cwchiu/MyTool
`, APP_VERSION),
}

func Execute(){
	rootCmd.Execute()
}