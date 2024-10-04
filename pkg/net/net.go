package net

import (
	"io"
	"net/http"
	"net/url"
)

// ValidateURL takes a URL and checks if it is valid.
// The function returns true if the URL is valid and false otherwise.
func ValidateURL(urlForValidate string) bool {
	_, err := url.ParseRequestURI(urlForValidate)
	return err == nil
}

// InitReader takes a URL and headers map and returns an io.ReadCloser, the content length of the response and an error.
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
