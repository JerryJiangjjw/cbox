package tools

import (
	"encoding/base64"
	"net/url"
	"time"
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

func TimeStampToDate(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

func DateToTimeStmap(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}
