package auth

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtClaims 自定义Claims
type JwtClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

// md5加密
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// sha256加密
func Sha256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// jwt 获取token , expire 为过期时间，单位为小时
func GenToken(uuid, salt string, expire int) (string, error) {
	expiration := time.Now().Add(time.Hour * time.Duration(expire))
	claims := JwtClaims{
		UUID: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			Issuer:    "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(salt))

	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func ParseToken(tokenString, salt string) (*JwtClaims, error) {
	claims := JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return &claims, nil
}
