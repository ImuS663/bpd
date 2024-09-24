package net

import (
	"io"
	"net/http"
	"net/url"
)

func ValidateURL(urlForValidate string) bool {
	_, err := url.ParseRequestURI(urlForValidate)
	return err == nil
}

func InitReader(url string, headers map[string]string) (io.ReadCloser, int64, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	setHeaders(req, headers)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	return resp.Body, resp.ContentLength, nil
}

func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
