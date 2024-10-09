package factory

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密函数
func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
