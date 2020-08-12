package ghttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func JSON(req *http.Request, v interface{}) error {
	client := http.Client{}
	// client := http.Client{Timeout: time.Second * 10} // TODO

	if req.URL.Host != "api.telegram.org" {
		log.Printf("JSON request: %v", req) // TODO: ?
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http.Client.Do error: %s", err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return fmt.Errorf("JSON Decoder error after \"%s %s\" (%s): %s", res.Proto, res.Status, req.Method, err)
	}

	return nil
}

func FormJSON(req *http.Request, data url.Values, v interface{}) error {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))

	return JSON(req, v)
}

func JSONJSON(req *http.Request, data interface{}, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")

	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %s", err)
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(b))

	return JSON(req, v)
}

func Success(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}

func NewAuthRequest(method, url, token string) (*http.Request, error) {
	req, err := NewRequest(method, url)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req, nil
}

// TODO: basic auth request

func NewRequest(method, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}
