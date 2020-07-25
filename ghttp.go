package ghttp

import (
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

func Success(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}
