package server

/**
 * 字元產生器 (CHARGEN) (Character Generator)
 *
 * 網路埠號
 * Port 19
 * 主要特性
 * 用戶端透過 Socket 建立連線到該服務，該服務會立刻回應永無止境的 ASCII 字元，直到用戶端主動發出斷線請求
 * 適用情境
 * 開發 Socket 應用程式時，可用來測試用戶端到伺服器端之間的「下載」頻寬大小，以及利用 Buffer 儲存大量文字的情境
 * 斷線時機
 * 用戶端主動發出斷線要求
 *
 * http://www.faqs.org/rfcs/rfc864.html
 */
import (
	"net"

	"github.com/spf13/cobra"
)

// SetupCharGenCommand 命令列入口
func SetupCharGenCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "chargen",
		Short: "CharGen Server",
		Long:  `CharGen Server`,
		Run: func(cmd *cobra.Command, args []string) {
			factory(13, func(conn net.Conn) {
				defer conn.Close()
				var s []byte
				for i := 33; i <= 125; i++ {
					s = append(s, byte(i))
				}

				for {
					conn.Write(s)
				}
			})
		},
	}

	rootCmd.AddCommand(cmd)
}
