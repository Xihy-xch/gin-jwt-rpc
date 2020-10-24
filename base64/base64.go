package base64

import "encoding/base64"

func EncodeToBase64(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}
func DecodeFromBase64(str string) string {
	result, _ := base64.URLEncoding.DecodeString(str)
	return string(result)
}
