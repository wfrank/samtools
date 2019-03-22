package isam

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Client struct {
	Host       string
	User       string
	Pass       string
	httpClient *http.Client
}

func NewClient(host, user, pass string) *Client {

	return &Client{
		Host: host,
		User: user,
		Pass: pass,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 3 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}
