package web

import (
	"fmt"
    "encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)


// {"ip":"223.139.225.222","address":"\u4e2d\u56fd\u53f0\u6e7e\u7701","state":"\u4e2d\u56fd","province":"\u53f0\u6e7e\u7701","city":"","district":"","longitude":120.36132599,"latitude":23.77722277,"status":0}
type JsonIPInfo struct {
    IP        string `json:"ip"`
    Address        string `json:"address"`
    State        string `json:"state"`
    Province        string `json:"province"`
    City        string `json:"city"`
    District        string `json:"district"`
    Longitude        float32 `json:"longitude"`
    Latitude        float32 `json:"latitude"`
    Status        int `json:"status"`
}

func SetupIPInfoCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "ipinfo <ip>",
		Short: "get ip info",
		Long:  `get ip info`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("required <ip>")
			}
			request := gorequest.New()
            ip := args[0]
			_, body, err := request.Get("http://www.bunian.cn/gongjv/ip/").Param("ip", ip).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
            // fmt.Println(body)
            var result JsonIPInfo
            json.Unmarshal([]byte(body), &result)
			// fmt.Println(result)
            if result.Status != 0 {
                fmt.Println("fail")
                return
            }
            fmt.Printf("State: %v\n", result.State)
            fmt.Printf("Province: %v\n", result.Province)
            fmt.Printf("City: %v\n", result.City)
            fmt.Printf("Address: %v\n", result.Address)
            fmt.Printf("Lat: %v, Lng: %v\n", result.Longitude, result.Longitude)
            
		},
	}
	rootCmd.AddCommand(cmd)

}
