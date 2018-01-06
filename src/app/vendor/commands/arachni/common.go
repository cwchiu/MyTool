package arachni

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
)

type JsonResult = map[string]interface{}
type JsonOption = map[string]interface{}

type Arachni struct {
	Entry    string
	Username string
	Password string
}

func CreateArachni(entry, username, password string) *Arachni {
	return &Arachni{
		Entry:    entry,
		Username: username,
		Password: password,
	}
}

func (self Arachni) newRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	if self.Username != "" && self.Password != "" {
		request.SetBasicAuth(self.Username, self.Password)
	}

	return request
}

func (self Arachni) StartScan(url string, options *JsonOption) (*JsonResult, error) {
	m := map[string]interface{}{
		"url": url,
	}
	if options != nil {
		for k, v := range *options {
			m[k] = v
		}
	}

	mJson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	// fmt.Println(mJson)
	// contentReader := bytes.NewReader(mJson)
	request := self.newRequest()
	api_url := fmt.Sprintf("%s/scans", self.Entry)
	request.Post(api_url).Send(string(mJson))
	_, body, errs := request.End()
	if errs != nil {
		return nil, errors.New("request error")
	}
	// fmt.Println(body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (self Arachni) GetScanResult(task_id string) (*JsonResult, error) {
	request := self.newRequest()
	url := fmt.Sprintf("%s/scans/%s", self.Entry, task_id)
	_, body, errs := request.Get(url).End()
	if errs != nil {
		return nil, errors.New("request error")
	}

	// fmt.Println(resp)
	// fmt.Println(body)
	if strings.Contains(body, "Scan not found for token") {
		return nil, errors.New("token not found")
	}
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}

	// for {
	// if data["status"].(string) == "done" {
	// break
	// }
	// fmt.Println("wait")
	// time.Sleep(1000 * time.Millisecond)
	// }

	return &data, nil
}
