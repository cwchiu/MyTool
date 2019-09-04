package netcat

import (
	"github.com/spf13/cobra"
)

func SetupClientCommand(rootCmd *cobra.Command) {
	var port int
	var is_udp bool
	var target string
	cmd := &cobra.Command{
		Use:   "client",
		Short: "用戶端",
		Run: func(cmd *cobra.Command, args []string) {
			if is_udp {
				udp_connect(target, port)
			} else {
				tcp_connect(target, port)
			}
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 1234, "監聽的通訊埠")
	cmd.Flags().BoolVarP(&is_udp, "udp", "u", false, "UDP 模式?")
	cmd.Flags().StringVarP(&target, "target", "t", "127.0.0.1", "要連線的主機")
	rootCmd.AddCommand(cmd)
}
