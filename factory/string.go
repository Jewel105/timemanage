package factory

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "0987654321"

// 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	lengh := len(charset)
	for i := range result {
		result[i] = charset[seededRand.Intn(lengh)]
	}
	return string(result)
}
