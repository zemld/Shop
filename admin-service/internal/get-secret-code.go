package internal

import "encoding/base64"

func GetSecretCode(admin string) string {
	return base64.StdEncoding.EncodeToString([]byte(admin + "secret"))
}
