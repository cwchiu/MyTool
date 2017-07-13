package server

/**
 * 回應服務 (Echo)
 *
 * 網路埠號
 * Port 7
 * 主要特性
 * 用戶端透過 Socket 傳送任何字串到該服務，該服務會立刻回應與用戶端輸入字串一模一樣的內容到用戶端
 * 斷線時機
 * 用戶端主動發出斷線要求
 * 適用情境
 * 可用來測試網路是否正常連線，也可用來監控網路是否斷線。
 * 開發 Socket 應用程式時，可用來測試用戶端與伺服器端彼此互動的情況
 */
import (
	"github.com/spf13/cobra"
	"log"
	"net"
)

func SetupEchoCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "echo",
		Short: "Echo Server",
		Long:  `Echo Server`,
		Run: func(cmd *cobra.Command, args []string) {
			Factory(7, func(conn net.Conn) {
				defer conn.Close()
				var buf = make([]byte, 10)
				var size = 0
				for {
					// read from the connection
					log.Println("start to read from conn")
					n, err := conn.Read(buf)
					if err != nil {
						log.Println("conn read error:", err)
						break
					}
					conn.Write(buf)
					log.Println(n)
					size += n
				}
				log.Printf("read %d bytes\n", size)
			})
		},
	}

	rootCmd.AddCommand(cmd)
}
