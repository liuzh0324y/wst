package wsthttp

import (
	"io"
	"log"
	"net/http"
)

// Put define http Put
func handlePut(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	c, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	// req.Header.Add("Content-Length", strconv.FormatInt(size, 10))

	log.Println("header info: \n", req.Header)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
