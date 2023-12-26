package client

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

func NewRequest(body map[string]interface{}, path string) *resty.Request {
	connection := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetHeader("Content-Type", "application/json").
		R().SetOutput(path).
		SetBody(&body)

	return connection
}
