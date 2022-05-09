package helper

import (
	"book-api/config"
	"book-api/model/web"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJWTToken(config config.JWTConfig, claims web.JWTCustomClaims) (string, error) {

	claims2 := &web.JWTCustomClaims{
		Id:    claims.Id,
		Email: claims.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)

	t, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil

}
