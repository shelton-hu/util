package netutil

import "net/url"

// 组装url
func BuildUrl(u string, p map[string]string) (string, error) {
	rawUrl, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	rawUrl.RawQuery = BuildQuery(p)
	return rawUrl.String(), nil
}

// 组装query参数
func BuildQuery(p map[string]string) string {
	params := url.Values{}
	for k, v := range p {
		params.Set(k, v)
	}
	return params.Encode()
}
