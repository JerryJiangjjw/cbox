package tools

import (
	"encoding/base64"
	"net/url"
)

func UrlEncode(urlStr string) string {
	urlencode := url.QueryEscape(urlStr)
	return urlencode
}

func UrlDecode(urlStr string) string {
	urldecode, _ := url.QueryUnescape(urlStr)
	return urldecode
}

func Base64Encode(b64Str string) string {
	base64encode := base64.StdEncoding.EncodeToString([]byte(b64Str))
	return base64encode
}

func Base64Decode(b64Str string) string {
	base64decode, _ := base64.StdEncoding.DecodeString(b64Str)
	return string(base64decode)
}
