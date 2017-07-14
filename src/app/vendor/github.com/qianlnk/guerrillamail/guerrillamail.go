package guerrillamail

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
    "log"
)

type GuerrillamailClient struct {
	SidToken   string
	apiKey     string
	apiToken   string
    debug      bool
	mailClient *http.Client
}

func NewGuerrillamailClient(cli *http.Client) *GuerrillamailClient {
	if cli == nil {
		cli = http.DefaultClient
	}

	return &GuerrillamailClient{
		mailClient: cli,
	}
}

func (g *GuerrillamailClient) SetDebug(enable bool)  {
    g.debug = enable
}

//Authorize Derive a token by calling a HMAC function
func (g *GuerrillamailClient) Authorize(apiKey string) error {
	//doesn't need authorize
	g.apiKey = apiKey
	if g.SidToken == "" || g.apiKey == "" || g.apiToken != "" {
		return nil
	}

	key := []byte(g.apiKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(g.SidToken))
	g.apiToken = base64.StdEncoding.EncodeToString(h.Sum(nil))

	return nil
}

func (g *GuerrillamailClient) apiURL(f string, args Argument) string {
	uri, err := url.Parse(GuerrillamailAPI)
	if err != nil {
		return ""
	}

	q := uri.Query()
	q.Set("f", f)

	if _, ok := args["sid_token"]; !ok && g.SidToken != "" {
		args["sid_token"] = g.SidToken
	}

	for k, v := range args {
		if k != "" && v != "" {
			q.Set(k, v)
		}
	}

	uri.RawQuery = q.Encode()
    if g.debug {
        log.Println( uri )
    }
	return uri.String()
}

func (g *GuerrillamailClient) getAPI(f string, args Argument, v interface{}) error {
	req, err := http.NewRequest("GET", g.apiURL(f, args), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Context-Type", "application/json")

	if g.apiToken != "" {
		req.Header.Set("Authorization", "ApiToken "+g.apiToken)
	}

	resp, err := g.mailClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) == "false" {
		return fmt.Errorf("%s false", f)
	}

    if g.debug {
        log.Println( string(body) )
    }
	return json.Unmarshal(body, v)
}

func (g *GuerrillamailClient) GetEmailAddress(args Argument) (*GetEmailAddressResponse, error) {
	var resp GetEmailAddressResponse

	err := g.getAPI("get_email_address", args, &resp)
	g.SidToken = resp.SidToken

	return &resp, err
}

func (g *GuerrillamailClient) SetEmailUser(args Argument) (*SetEmailUserResponse, error) {
	var resp SetEmailUserResponse
	err := g.getAPI("set_email_user", args, &resp)
	g.SidToken = resp.SidToken

	return &resp, err
}

func (g *GuerrillamailClient) CheckEmail(args Argument) (resp *CheckEmailResponse, err error) {
	err = g.getAPI("check_email", args, &resp)

	return
}

func (g *GuerrillamailClient) DelEmail(args Argument) (resp *DelEmailResponse, err error) {
	err = g.getAPI("del_email", args, &resp)

	return
}

func (g *GuerrillamailClient) GetEmailList(args Argument) (resp *GetEmailListResponse, err error) {
	err = g.getAPI("get_email_list", args, &resp)

	return
}

func (g *GuerrillamailClient) FetchEmail(args Argument) (resp *FetchEmailResponse, err error) {
	err = g.getAPI("fetch_email", args, &resp)

	return
}

func (g *GuerrillamailClient) ForgetMe(args Argument) (resp *FetchEmailResponse, err error) {
	err = g.getAPI("forget_me", args, &resp)

	return
}
