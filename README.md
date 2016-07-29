# sendcloud [![GoDoc](https://godoc.org/github.com/miaolz123/sendcloud?status.svg)](https://godoc.org/github.com/miaolz123/sendcloud) [![Build Status](https://travis-ci.org/miaolz123/sendcloud.svg?branch=master)](https://travis-ci.org/miaolz123/sendcloud)

### A SDK of sendcloud.net for Golang

```go
package main

import (
	"log"

	"github.com/miaolz123/sendcloud"
)

// Please replace this section
const (
	emailAPIUser = "balabala"
	emailAPIKey  = "balabalabalabala"
    sendAddr     = "GooGle Inc.<test@balabala.sendcloud.org>"
    receiveAddr  = "yourname@126.com"
)

func main() {
    conf := sendcloud.Config{
		EmailAPIUser: emailAPIKey,
		EmailAPIKey:  emailAPIKey,
	}
	client := sendcloud.New(conf)
	xs := sendcloud.XSmtpAPI{
		To: []string{receiveAddr},
		Sub: map[string][]string{
			"%name%": []string{"TEST USER"},
		},
	}
	if err := client.SendEmailTpl(sendAddr, "test_template_active", xs); err != nil {
		log.Println("Send email error:", err)
	}
}
```
