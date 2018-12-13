package wsthttp

import "net/http"

// Get overload http.Get
func handleGet(url string) (*http.Response, error) {
	c, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	return c.Get(url)
}
