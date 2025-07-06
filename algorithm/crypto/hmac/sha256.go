package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
)

func Sha256Bytes(s, key string) []byte {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return hashed.Sum(nil)
}

func Sha256Str(s, key string) string {
	return string(Sha256Bytes(s, key))
}
