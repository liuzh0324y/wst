package wsthttp_test

import (
	"strings"
	"testing"

	"github.com/wsterr/wst/wsthttp"
)

const (
	defaultURL = "https://llvision.com"
)

func TestHttpGet(t *testing.T) {
	resp, err := wsthttp.Get(defaultURL)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Body)
	}
}

func TestHttpPut(t *testing.T) {
	resp, err := wsthttp.Put(defaultURL, "application/json", strings.NewReader("wsthttp put."))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Body)
	}
}

func TestHttpPost(t *testing.T) {
	resp, err := wsthttp.Post(defaultURL, "application/json", strings.NewReader("wsthttp post."))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Body)
	}
}
