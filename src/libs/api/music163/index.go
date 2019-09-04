package music163

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"github.com/cwchiu/MyTool/libs"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	nonce   = "0CoJUm6Qyw8W8jud"
	pubKey  = "010001"
	keys    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
	iv      = "0102030405060708"
)

var userAgentList = [19]string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows NT 6.3; Win64, x64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
}

type SongInfo struct {
	Url     string `json:"url"`
	Size    int    `json:"size"`
	BitRate int    `json:"br"`
	Type    string `json:"type"`
}

type QuerySongResponse struct {
	Code int        `json:"code"`
	Data []SongInfo `json:"data"`
}

type MainSong struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProgramInfo struct {
	Song MainSong `json:"mainSong"`
}

type QueryProgramResponse struct {
	Code    int         `json:"code"`
	Program ProgramInfo `json:"program"`
}

func DownloadSong(sid, filename string) error {
	songs := GetDownloadLink(sid)
	if len(songs) == 0 {
		return errors.New("download link not found")
	}
	for _, bs := range []int{320000, 192000, 128000} {
		song, ok := songs[bs]
		if ok {
			fmt.Println(song.Url)
			err := libs.DownloadOne(song.Url, filename)
			if err == nil {
				return nil
			}
			fmt.Printf("%s: %s\n", song.Url, err)
		}
	}

	return errors.New("download fail")
}

func GetDownloadLink(sid string) map[int]SongInfo {
	chs := make(chan *QuerySongResponse, 3)
	defer func() {
		close(chs)

	}()
	go QueryDownloadLink(sid, "320000", chs)
	go QueryDownloadLink(sid, "192000", chs)
	go QueryDownloadLink(sid, "128000", chs)

	ret := map[int]SongInfo{}
	for i := 0; i < 3; i += 1 {
		resp := <-chs
		if resp != nil {
			si := (*resp).Data[0]
			ret[si.BitRate] = si
		}
	}
	return ret
}

func QueryDownloadLink(ids, rate string, ch chan *QuerySongResponse) {
	preParams := fmt.Sprintf(`{"ids": "[%s]", "br": %s, "csrf_token": ""}`, ids, rate)
	body, err := Music163Encrypt(preParams)
	if err != nil {
		ch <- nil
		return
	}
	res, err := post("http://music.163.com/weapi/song/enhance/player/url?csrf_token=", body)
	if err != nil {
		ch <- nil
		return
	}
	var tmp QuerySongResponse
	err = json.Unmarshal(res, &tmp)
	if err != nil {
		ch <- nil
		return
	}

	ch <- &tmp
}

func DownloadRadio(id string) error {
	resp, err := QueryRadioInfo(id)
	if err != nil {
		return err
	}

	if (*resp).Code != 200 {
		return errors.New("get radio fail")
	}

	song := (*resp).Program.Song
	return DownloadSong(strconv.Itoa(song.Id), song.Name+".mp3")
}

// 電台節目音樂連結查詢
func QueryRadioInfo(id string) (*QueryProgramResponse, error) {
	preParams := fmt.Sprintf(`{csrf_token:"",id:%s}`, id)
	body, err := Music163Encrypt(preParams)
	if err != nil {
		return nil, err
	}
	// fmt.Println(preParams)
	// fmt.Println(body)
	res, err := post("http://music.163.com/weapi/dj/program/detail?csrf_token=", body)
	if err != nil {
		return nil, err
	}

	var tmp QueryProgramResponse
	err = json.Unmarshal(res, &tmp)
	if err != nil {
		return nil, err
	}
	return &tmp, nil
}

// 通过 CBC模式的AES加密 用 sKey 加密 sSrc
func aesEncrypt(sSrc string, sKey string) (string, error) {
	iv := []byte(iv)
	block, err := aes.NewCipher([]byte(sKey))
	if err != nil {
		return "", err
	}
	padding := block.BlockSize() - len([]byte(sSrc))%block.BlockSize()
	src := append([]byte(sSrc), bytes.Repeat([]byte{byte(padding)}, padding)...)
	model := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(src))
	model.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func post(api_url, params string) ([]byte, error) {
	request, err := http.NewRequest("POST", api_url, bytes.NewBuffer([]byte(params)))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("User-Agent", fakeAgent())

	proxyUrl, err := url.Parse("http://proxy.uku.im:443")
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}

	client := &http.Client{}
	response, reqErr := client.Do(request)
	if reqErr != nil {
		return nil, reqErr
	}
	defer response.Body.Close()
	resBody, resErr := ioutil.ReadAll(response.Body)
	if resErr != nil {
		return nil, resErr
	}
	return resBody, nil
}

func fakeAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentList[r.Intn(19)]
}

func Music163Encrypt(sSrc string) (string, error) {
	// expect := "M+sMQRArh9OdctOmgLB+lcKAiWtieKy6+rcGKYzo8t9JXhGHXdUtpVgsRHsgiVyeDbtkFD8y3lsFtKgXoWVGVA=="
	actual, _ := aesEncrypt(sSrc, "0CoJUm6Qyw8W8jud")
	actual2, _ := aesEncrypt(actual, "a8LWv2uAtXjzSfkQ")
	actual3 := url.QueryEscape(actual2)
	ret := fmt.Sprintf("params=%s&encSecKey=2d48fd9fb8e58bc9c1f14a7bda1b8e49a3520a67a2300a1f73766caee29f2411c5350bceb15ed196ca963d6a6d0b61f3734f0a0f4a172ad853f16dd06018bc5ca8fb640eaa8decd1cd41f66e166cea7a3023bd63960e656ec97751cfc7ce08d943928e9db9b35400ff3d138bda1ab511a06fbee75585191cabe0e6e63f7350d6", actual3)
	return ret, nil
}
