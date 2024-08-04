package jwtutil

import (
	"Gin_DEMO/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成token
var jwtSignKey = []byte(config.JwtSignKey)

type MyCustomClaims struct {
	Username string `json:"foo"`
	jwt.RegisteredClaims
}

// 解析token
func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtEepireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "dotbalo",
			Subject:   "oufan",
		},
	}
	//生成token:选择生成算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	return ss, err
}
