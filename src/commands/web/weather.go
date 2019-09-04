package web

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type weather_data struct {
	ShiDu   string `json:"shidu"`
	WenDu   string `json:"wendu"`
	Quality string `json:"quality"`
}

type response struct {
	Message string        `json:"message"`
	Data    *weather_data `json:"data"`
}

func SetupWeatherCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "weather <city>",
		Short: "查詢城市天氣",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("required <number>")
			}

			request := getRequest(false, false)
			_, body, err := request.Get("http://www.sojson.com/open/api/weather/json.shtml").Param("city", args[0]).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			var result response
			json.Unmarshal([]byte(body), &result)
			if result.Message != "Success !" {
				fmt.Println(result.Message)
				return
			}

			fmt.Println("濕度:", result.Data.ShiDu)
			fmt.Println("溫度:", result.Data.WenDu)
			fmt.Println("空氣品質:", result.Data.Quality)
		},
	}
	rootCmd.AddCommand(cmd)

}
