package common

import (
	"Kjasn/ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 发放 token， 返回 token 字符串和 错误信息
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 过期时间 7d

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Kjasn",
			Subject:   "user token",
		},
	}

	// 使用HS256算法进行token生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 从 tokenString 中解析出 claims 信息,  返回 token 对象，claims 信息， 错误信息
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	// 解析出 token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}

// 以下是某一次登录获得的 token

// 由三部分组成  通过 . 分开  第一部分是 header，保存的信息是加密协议， 第二部分是 payload 第三部分是 signature
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
// eyJVc2VySWQiOjQsImV4cCI6MTY5NzU1MzA1OCwiaWF0IjoxNjk2OTQ4MjU4LCJpc3MiOiJLamFzbiIsInN1YiI6InVzZXIgdG9rZW4ifQ.
// SVxoFAB9U1w_XmyXbBQHpUgrcTVcoRImRV_9I9WvmdA
