package wsthttp

import (
	"crypto/tls"
	"io"
	"net/http"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

var wstClient = &Client{
	client: http.Client{
		Transport: tr,
	},
}

// Client  define http client
type Client struct {
	client http.Client
}

// New Create a http instance
func New(useTLS bool) *Client {
	var client http.Client
	if useTLS != true {
		client = http.Client{}
	} else {
		client = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}

	return &Client{client: client}
}

func Get(url string) (*http.Response, error) {
	return handleGet(url)
}

func Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return handlePost(url, contentType, body)
}

func Put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return handlePut(url, contentType, body)
}
