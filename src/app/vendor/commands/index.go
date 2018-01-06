package commands

import (
	"github.com/spf13/cobra"
)

var	rootCmd = &cobra.Command{Use: "tool", Long: `我的個人常用工具
Site: https://chuiwenchiu.wordpress.com
Github: https://github.com/cwchiu/MyTool
`}

func Execute(){
	rootCmd.Execute()
}