package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtOption func(*Jwt)

func JwtSigningMethod(method jwt.SigningMethod) JwtOption {
	return func(j *Jwt) {
		j.SigningMethod = method
	}
}

func JwtSecretKey(secret string) JwtOption {
	return func(j *Jwt) {
		j.SecretKey = []byte(secret)
	}
}

func JwtExpire(expire time.Duration) JwtOption {
	return func(j *Jwt) {
		j.Expire = expire
	}
}

func JwtRefresh(refresh time.Duration) JwtOption {
	return func(j *Jwt) {
		j.Refresh = refresh
	}
}
