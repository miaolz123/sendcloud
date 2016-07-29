package sendcloud

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// XSmtpAPI : X-SMTPAPI for SDK
type XSmtpAPI struct {
	To      []string            `json:"to"`
	Sub     map[string][]string `json:"sub"`
	Section map[string]string   `json:"section,omitempty"`
}

// SendEmail : send a email
func (c Client) SendEmail(from, to, subject, html string) (err error) {
	params := url.Values{
		"apiUser": {c.config.EmailAPIUser},
		"apiKey":  {c.config.EmailAPIKey},
		"from":    {from},
		"to":      {to},
		"subject": {subject},
		"html":    {html},
	}
	body, err := c.httpDo("POST", emailURL+"mail/send", params)
	if err != nil {
		return err
	}
	result := struct {
		Result  bool
		Message string
	}{false, "未知错误"}
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if !result.Result {
		err = fmt.Errorf(result.Message)
	}
	return err
}

// SendEmailTpl : send a template email
func (c Client) SendEmailTpl(from, tplName string, xsmtpapi XSmtpAPI) (err error) {
	xs, _ := json.Marshal(xsmtpapi)
	params := url.Values{
		"apiUser":            {c.config.EmailAPIUser},
		"apiKey":             {c.config.EmailAPIKey},
		"from":               {from},
		"templateInvokeName": {tplName},
		"xsmtpapi":           {string(xs)},
	}
	body, err := c.httpDo("POST", emailURL+"mail/sendtemplate", params)
	if err != nil {
		return err
	}
	result := struct {
		Result  bool
		Message string
	}{false, "未知错误"}
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if !result.Result {
		err = fmt.Errorf(result.Message)
	}
	return err
}
