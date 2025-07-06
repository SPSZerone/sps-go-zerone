package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}
