package server

/**
 * 每日一句 (QUOTE)
 *
 * 網路埠號
 * Port 17
 * 主要特性
 * 用戶端透過 Socket 建立連線到該服務，該服務會立刻回應一段話 (通常是一句名人說過的話)，回應完畢後立即斷線
 * 這一句話，可能為一行或多行文字，但規定不得大於 512 位元組
 * 適用情境
 * 開發 Socket 應用程式時，可用來取得不固定大小的文字，或測試伺服器端主動發出斷線的狀況
 * 斷線時機
 * 伺服器端主動發出斷線要求
 * http://www.faqs.org/rfcs/rfc865.html
 */
import (
	"math/rand"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var data = []string{
	"Do not, for one repulse, forgo the purpose that you resolved to effort. ( Shakespeare ) ",
	"The man who has made up his mind to win will never say \" Impossible\".( Napoleon )",
	"Miracles sometimes occur, but one has to work terribly for them. ( C. Weizmann ) ",
	"There is no such thing as darkness; only a failure to see. ( Muggeridge ) ",
	"Time is a bird for ever on the wing. ( T. W. Robertson ) ",
	"If you do not learn to think when you are young, you may never learn. ( Edison ) ",
	"A day is a miniature of eternity. ( Emerson ) ",
	"Morality may consist solely in the courage of making a choice. ( L. Blum ) ",
	"If there were less sympathy in the world, there would be less trouble in the world. ( O. Wilde ) ",
}

func SetupQuoteCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "quote",
		Short: "Quote Server",
		Long:  `Quote Server`,
		Run: func(cmd *cobra.Command, args []string) {
			factory(17, func(conn net.Conn) {
				defer conn.Close()
				rand.Seed(time.Now().UnixNano())
				var s = data[rand.Intn(9)]
				conn.Write([]byte(s))
			})
		},
	}

	rootCmd.AddCommand(cmd)
}
