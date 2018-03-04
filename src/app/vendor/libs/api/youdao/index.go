package youdao

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	// "net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type TranslateResult struct {
	Source string `json:"src"`
	Target string `json:"tgt"`
}

type YoudaoTranslateResult struct {
	Type             string              `json:"type"`
	ErrorCode        int                 `json:"errorCode"`
	TranslateResults [][]TranslateResult `json:"translateResult"`
}

func Translate(word, to, from string) (string, error) {

	form := url.Values{}
	form.Add("i", word)
	form.Add("from", from)
	form.Add("to", to)
	form.Add("smartresult", "dict")
	form.Add("client", "fanyideskweb")
	form.Add("keyfrom", "fanyi.web")
	form.Add("doctype", "json")
	form.Add("version", "2.1")
	form.Add("action", "FY_BY_REALTIME")
	form.Add("typoResult", "false")
	salt := fmt.Sprintf("%d000", time.Now().UTC().Unix())
	form.Add("salt", salt)

	h := md5.New()
	h.Write([]byte("fanyideskweb" + word + salt + "ebSeFb%=XZ%T[KZ)c(sy!"))
	sign := hex.EncodeToString(h.Sum(nil))
	form.Add("sign", sign)
	data := form.Encode()
	// fmt.Println( data  )
	// data := `i=hel&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=1520166748938&sign=2c02a518d877205bfa9b214c7e3efe18&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTIME&typoResult=false`
	request, err := http.NewRequest("POST", "http://fanyi.youdao.com/translate?smartresult=dict&smartresult=rule", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Origin", "http://fanyi.youdao.com")
	request.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Referer", "http://fanyi.youdao.com/")

	// dump, _ := httputil.DumpRequestOut(request, true)
	// fmt.Println(string(dump))

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()

	// dump2, _ := httputil.DumpResponse(response, true)
	// fmt.Println(string(dump2))

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println(body)
	var result YoudaoTranslateResult
	json.Unmarshal([]byte(body), &result)
	// fmt.Println(result)
	if result.ErrorCode != 0 {
		return "", errors.New("translate fail")
	}

	// fmt.Println(result.Type)
	// fmt.Println(result)
	// fmt.Println(result.TranslateResults[0][0].Source)
	return result.TranslateResults[0][0].Target, nil
}

func DictQuery(word string) (string, error) {
	var qs url.Values
	qs.Add("keyfrom", "dict.index")
	qs.Add("q", word)
	url := fmt.Sprintf("http://dict.youdao.com/search?%s", qs.Encode())

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}
	// fmt.Println(body)
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return "", err
	}
	var result strings.Builder
	doc.Find("#results-contents > #phrsListTab > .trans-container li").Each(func(i int, s *goquery.Selection) {
		result.WriteString(s.Text() + "\n")
	})

	return result.String(), nil
}
