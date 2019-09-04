package github

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
    "bytes"
    "io/ioutil"
)

type GistFileContent struct {
	Content string `json:"content"`
}

type GistForm struct {
	Desc   string                     `json:"description"`
	Public bool                       `json:"public"`
	Files  map[string]GistFileContent `json:"files"`
}

type GistResponse struct {
	Url string `json:"html_url"`
}

type GistConfig struct {
	Title    string
	Content  string
	Desc     string
	Secret   bool
	Username string
	Token    string
}

func NewGistConfig() GistConfig {
	cfg := GistConfig{
		Secret: true,
		Title:  "Untitle",
	}

	return cfg
}

func CreateGist(params GistConfig) (*GistResponse, error) {

	data := GistForm{
		Desc:   params.Desc,
		Public: !params.Secret,
		Files: map[string]GistFileContent{
			params.Title: GistFileContent{
				Content: params.Content,
			},
		},
	}
	bs, err := json.Marshal(data)

	request, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")

	if params.Username != "" && params.Token != "" {
		request.Header.Set("User-Agent", "Awesome-Octocat-App")
		hash := base64.StdEncoding.EncodeToString([]byte(params.Username + ":" + params.Token))
		request.Header.Set("Authorization", "Basic "+hash)
	}
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(body)
	var result GistResponse
	err = json.Unmarshal([]byte(body), &result)
	// fmt.Println(result.Url)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
