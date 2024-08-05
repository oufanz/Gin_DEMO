package jwtutil

import (
	"Gin_DEMO/config"
	"Gin_DEMO/utils/logs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

type MyCustomClaims struct {
	Username string `json:"foo"`
	jwt.RegisteredClaims
}

// 生成token

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

// 解析token
func ParseToken(ss string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		logs.Error(nil, "解析失败")
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims, nil
	} else {
		logs.Error(nil, "Token不合法")
		return nil, errors.New("Token不合法")
	}

}
