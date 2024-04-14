package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenSign(body, secretKey string) string {
	// md5
	hash := md5.New()
	hash.Write([]byte(body + "." + secretKey))
	return hex.EncodeToString(hash.Sum(nil))
}
