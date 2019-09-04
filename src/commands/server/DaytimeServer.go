package server

/**
 *  日期時間服務 (Daytime)
 *
 *  網路埠號
 *  Port 13
 *  主要特性
 *  用戶端透過 Socket 建立連線到該服務，該服務會立刻回應伺服器上的目前時間，回應完畢後立即斷線
 *  適用情境
 *  可用來取得伺服器上目前的日期時間
 *  開發 Socket 應用程式時，可用來取得伺服器時間，或測試伺服器端主動發出斷線的狀況
 *  斷線時機
 *  伺服器端主動發出斷線要求
 *  測試方法
 *  透過 telnet 工具程式連上本機的 Port 13
 *  當看見伺服器回應的日期時間字串後會立即斷線
 */
import (
	"github.com/spf13/cobra"
	"net"
	"time"
)

// SetupDaytimeCommand CLI Entrypoint
func SetupDaytimeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "daytime",
		Short: "DatTime Server",
		Long:  `DatTime Server`,
		Run: func(cmd *cobra.Command, args []string) {
			factory(13, func(conn net.Conn) {
				defer conn.Close()
				now := time.Now()
				dateFmt := "2006/01/02 15:04:05"
				conn.Write([]byte(now.UTC().Format(dateFmt)))
			})
		},
	}

	rootCmd.AddCommand(cmd)
}
