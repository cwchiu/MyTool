package web

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

// {"ip":"223.139.225.222","address":"\u4e2d\u56fd\u53f0\u6e7e\u7701","state":"\u4e2d\u56fd","province":"\u53f0\u6e7e\u7701","city":"","district":"","longitude":120.36132599,"latitude":23.77722277,"status":0}
type bunianIPInfo struct {
	IP        string  `json:"ip"`
	Address   string  `json:"address"`
	State     string  `json:"state"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Status    int     `json:"status"`
}

// {"code":0,"data":{"country":"\u7f8e\u56fd","country_id":"US","area":"","area_id":"","region":"","region_id":"","city":"","city_id":"","county":"","county_id":"","isp":"","isp_id":"","ip":"8.8.8.8"}}
type taobaoIPInfoData struct {
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	Area      string `json:"area"`
	AreaID    string `json:"area_id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	County    string `json:"county"`
	CountyID  string `json:"county_id"`
	ISP       string `json:"isp"`
	ISPID     string `json:"isp_id"`
	Ip        string `json:"ip"`
}

type taobaoIPInfo struct {
	Code int              `json:"code"`
	Data taobaoIPInfoData `json:"data"`
}
type taobaoIPInfoError struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func bunianQuery(ip string) {
	request := gorequest.New()
	_, body, err := request.Get("http://www.bunian.cn/gongjv/ip/").Param("ip", ip).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
	if err != nil {
		panic(err)
	}
	var result bunianIPInfo
	json.Unmarshal([]byte(body), &result)
	if result.Status != 0 {
		fmt.Println("fail")
		return
	}
	fmt.Printf("State: %v\n", result.State)
	fmt.Printf("Province: %v\n", result.Province)
	fmt.Printf("City: %v\n", result.City)
	fmt.Printf("Address: %v\n", result.Address)
	fmt.Printf("Lat: %v, Lng: %v\n", result.Longitude, result.Longitude)
}

func taobaoQuery(ip string) {
	request := gorequest.New()
	_, body, err := request.Get("http://ip.taobao.com/service/getIpInfo.php").Param("ip", ip).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
	if err != nil {
		panic(err)
	}
	// fmt.Println(body)
	var result taobaoIPInfo
	err_json := json.Unmarshal([]byte(body), &result)
	if err_json != nil {
		var result2 taobaoIPInfoError
		err_json = json.Unmarshal([]byte(body), &result2)
		if err_json != nil {
			panic(err_json)
		} else {
			fmt.Println(result2.Data)
			return
		}
	}
	data := result.Data

	fmt.Printf("Country: %v(%v)\n", data.Country, data.CountryID)
	fmt.Printf("Area: %v(%v)\n", data.Area, data.AreaID)
	fmt.Printf("ISP: %v(%v)\n", data.ISP, data.ISPID)
	fmt.Printf("Region: %v(%v)\n", data.Region, data.RegionID)
	fmt.Printf("City: %v(%v)\n", data.City, data.CityID)
}

func SetupIPInfoCommand(rootCmd *cobra.Command) {
	var src string
	cmd := &cobra.Command{
		Use:   "ipinfo <ip>",
		Short: "取得 IP 相關資訊",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <ip>")
			}
			if src == "taobao" {
				taobaoQuery(args[0])
			} else {
				bunianQuery(args[0])
			}
		},
	}

	cmd.Flags().StringVarP(&src, "src", "s", "taobao", "服務供應商: taobao, bunian, ")
	rootCmd.AddCommand(cmd)

}
