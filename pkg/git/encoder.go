package git

import "encoding/base64"

func EncodeToken(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(token))
}

func DecodeToken(token string) string {
	data, _ := base64.StdEncoding.DecodeString(token)
	return string(data)

}
