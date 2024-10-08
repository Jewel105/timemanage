package util

import (
	"gin_study/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	TokenId  uuid.UUID
	UserID   uint
	UserName string
	jwt.RegisteredClaims
}

var jwtSecret = []byte(config.JwtSecret)

func CreateToken(username string, userID uint) (string, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	claims := UserClaims{
		tokenId,
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // 过期时间3天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(jwtSecret))
	return s, err
}

func DecodeToken(tokenstring string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
