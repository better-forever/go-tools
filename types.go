package utils

import "github.com/dgrijalva/jwt-go"

// JwtClaims 自定义Claims
type JwtClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}
