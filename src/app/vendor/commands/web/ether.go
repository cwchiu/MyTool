package web

import (
	"fmt"
	"net/url"
    "encoding/json"
	"github.com/gorilla/websocket"
    "time"
	"github.com/spf13/cobra"
)


func api(){
    u := url.URL{Scheme: "wss", Host: "api.bitfinex.com", Path: "/ws"}
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    defer c.Close()
    
    done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				panic(err)
			}
			// fmt.Printf("recv: %s", message)
            var datas []float64
            err = json.Unmarshal([]byte(message), &datas)
            if err == nil {
                if len(datas)>2 {
                    fmt.Println(datas[1])
                }
            }
		}
	}()
    
    err = c.WriteMessage(websocket.TextMessage, []byte(`{"event":"subscribe","channel":"ticker","pair":"ETHUSD"}`))
    if err != nil {
        panic(err)
    }
    
    ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
        }
    }
}

func SetupEtherCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "ether",
		Short: "目前以太幣價格",
		Run: func(cmd *cobra.Command, args []string) {
			api()
		},
	}
	rootCmd.AddCommand(cmd)

}
