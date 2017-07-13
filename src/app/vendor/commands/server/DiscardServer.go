package server

/**
 * 丟棄服務 (Discard)
 *
 * 網路埠號
 * Port 9
 * 主要特性
 * 用戶端透過 Socket 傳送任何字串到該服務，該服務會立刻丟棄所接收到的任何字串
 * 適用情境
 * 可用來測試網路用戶端是否能夠正常透過 Socket 傳送資料
 * 你可以從用戶端大量的傳送資料，直到用戶端主動斷線為止
 * 開發 Socket 應用程式時，可用來測試用戶端到伺服器端之間的「上傳」頻寬大小
 * 斷線時機
 * 用戶端主動發出斷線要求
 */
import (
	"github.com/spf13/cobra"
	"log"
	"net"
)

func SetupDiscardCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "discard",
		Short: "Discard Server",
		Long:  `Discard Server`,
		Run: func(cmd *cobra.Command, args []string) {
			Factory(9, func(conn net.Conn) {
				defer conn.Close()
				var buf = make([]byte, 1024)
				for {
					n, err := conn.Read(buf)
					if err != nil {
						log.Println("conn read error:", err)
						break
					}
					log.Println(n)
				}
			})
		},
	}

	rootCmd.AddCommand(cmd)
}
