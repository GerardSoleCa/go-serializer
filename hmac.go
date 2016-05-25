package serializer

import (
	"encoding/base64"
	"strings"
	"crypto/sha1"
	"crypto/hmac"
)

func signStr(str string, key string) string {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write([]byte(strings.TrimSpace(str)))
	result := hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(result)
}