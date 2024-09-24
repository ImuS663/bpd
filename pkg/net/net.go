package net

import "net/url"

func ValidateURL(urlForValidate string) bool {
	_, err := url.ParseRequestURI(urlForValidate)
	return err == nil
}
