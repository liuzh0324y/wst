package wsthttp

import (
	"io"
	"net/http"
)

// Post overload http.Post
func handlePost(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	c, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	return c.Post(url, contentType, body)
}
