package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenJwtToken(accessExpire int64, accessSecret string, payloads map[string]interface{}) (string, error) {

	now := time.Now().Unix()

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(accessSecret))
}
