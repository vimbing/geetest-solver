package utils

import "net/url"

func Params(input map[string]string) string {
	values := url.Values{}

	for k, v := range input {
		values.Set(k, v)
	}

	return values.Encode()
}
