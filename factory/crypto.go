package factory

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5加密函数
func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA-256哈希函数封装
func Sha256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
