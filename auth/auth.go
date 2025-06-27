package auth

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
		uuid,
		jwt.StandardClaims{
			Issuer:    "bot-subscribe",
			ExpiresAt: expiration.Unix(),
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
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("伪造的token")
		}
		return nil, fmt.Errorf("无效的token")
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的token")
}
