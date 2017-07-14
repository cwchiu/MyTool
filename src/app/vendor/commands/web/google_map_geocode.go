package web

import (
	"fmt"
    "encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type Location struct {
    Lat float32 `json:"lat"`
    Lng float32 `json:"lng"`
}

type Geometry struct {
    Location Location `json:"location"`
}

type GeoResult struct {
    Geometry Geometry `json:"geometry"`
}

type GeoResults struct {
    Results       []GeoResult `json:"results"`
    Status        string `json:"status"`
}
func SetupGoogleMapGeocodeCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "geo <address>",
		Short: "address to lat and lng",
		Long:  `address to lat and lng`,
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
				panic("required <address>")
			}
			request := gorequest.New()
            address := args[0]
			_, body, err := request.Get("http://maps.google.com/maps/api/geocode/json").Param("sensor","false").Param("address", address).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
            // fmt.Println(body)
            var result GeoResults
            json.Unmarshal([]byte(body), &result)
			// fmt.Println(result)
            if result.Status != "OK" {
                fmt.Println("fail")
                return
            }
            if len(result.Results) < 1 {
                fmt.Println("not found")
                return
            }
            loc := result.Results[0].Geometry.Location
            fmt.Printf("Lat: %v, Lng: %v", loc.Lat, loc.Lng)
            
		},
	}
	rootCmd.AddCommand(cmd)

}
