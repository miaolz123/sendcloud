package sendcloud

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	emailURL = "http://api.sendcloud.net/apiv2/"
)

// Config : the config of SDK
type Config struct {
	EmailAPIUser string
	EmailAPIKey  string
	SmsAPIUser   string
	SmsAPIKey    string
}

// Client : a SDK client for sendcloud.net
type Client struct {
	client *http.Client
	config Config
}

// New : make a SDK client
func New(conf Config) Client {
	return Client{&http.Client{}, conf}
}

func (c Client) httpDo(method, url string, params url.Values) (body []byte, err error) {
	req, err := http.NewRequest(method, url, strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
