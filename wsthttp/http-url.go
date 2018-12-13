package wsthttp

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

func parseURL(urlname string) (*http.Client, error) {
	var c http.Client
	u, err := url.Parse(urlname)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "https" {
		c = http.Client{}
	} else {
		c = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}

	return &c, nil
}
