package whoscall

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type QueryResult struct {
	Name string
	Info string
}

func Query(number string) (*QueryResult, error) {
	url := fmt.Sprintf("https://whoscall.com/zh-TW/tw/%s/", number)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// fmt.Println(body)
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	var result QueryResult
	result.Name = strings.TrimSpace(doc.Find(".number-info .number-info__name").First().Text())
	result.Info = strings.TrimSpace(doc.Find(".number-info .number-info__subinfo").First().Text())
	return &result, nil
}
