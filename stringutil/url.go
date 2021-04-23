package stringutil

import (
	"net/url"
)

// url解码
func UrlDecoder(encodedUrl string) (string, error) {
	decodedUrl, err := url.QueryUnescape(encodedUrl)
	if err != nil {
		return "", err
	}
	return decodedUrl, nil
}

// url编码
func UrlEncoder(decodedUrl string) string {
	encodedUrl := url.QueryEscape(decodedUrl)
	return encodedUrl
}
