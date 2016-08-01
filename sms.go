package sendcloud

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
)

// SendSMS : send a text SMS
func (c Client) SendSMS(phone, templateID string, vars map[string]string) (err error) {
	varses, _ := json.Marshal(vars)
	params := url.Values{
		"phone":      {phone},
		"smsUser":    {c.config.SmsAPIUser},
		"templateId": {templateID},
	}
	signature := c.sign(params, string(varses))
	params.Add("vars", string(varses))
	params.Add("signature", signature)
	body, err := c.httpDo("POST", smsURL+"send", params)
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

// SendSMSVoice : send a voice SMS
func (c Client) SendSMSVoice(phone, code string) (err error) {
	params := url.Values{
		"code":    {code},
		"phone":   {phone},
		"smsUser": {c.config.SmsAPIUser},
	}
	signature := c.sign(params, "")
	params.Add("signature", signature)
	body, err := c.httpDo("POST", smsURL+"sendVoice", params)
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

func (c Client) sign(params url.Values, vars string) string {
	paramStr := params.Encode()
	if vars != "" {
		paramStr += "&vars=" + vars
	}
	paramStr = c.config.SmsAPIKey + "&" + paramStr + "&" + c.config.SmsAPIKey
	m := md5.New()
	m.Write([]byte(paramStr))
	return hex.EncodeToString(m.Sum(nil))
}
