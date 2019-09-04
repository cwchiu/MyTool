package netcat

import (
	"github.com/spf13/cobra"
)

func SetupServerCommand(rootCmd *cobra.Command) {
	var port int
	var is_udp bool
	cmd := &cobra.Command{
		Use:   "server",
		Short: "伺服器",
		Run: func(cmd *cobra.Command, args []string) {
			if is_udp {
				udp_listen(port)
			} else {
				tcp_listen(port)
			}
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 1234, "監聽的通訊埠")
	cmd.Flags().BoolVarP(&is_udp, "udp", "u", false, "UDP 模式?")
	rootCmd.AddCommand(cmd)
}
