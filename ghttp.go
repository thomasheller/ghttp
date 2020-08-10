package ghttp

import (
	"bytes"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func JSON(req *http.Request, v interface{}) error {
	client := http.Client{}
	// client := http.Client{Timeout: time.Second * 10} // TODO

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(v)
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
		return err
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

