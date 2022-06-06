package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("seckilltode")

type Claims struct {
	Uid uint `json:"uid"`
	jwt.StandardClaims
}

//签发用户
func GenerateToken(uid uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	Claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "seckill",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//验证用户token UID
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
