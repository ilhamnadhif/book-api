package web

import "github.com/golang-jwt/jwt"

type (
	TokenResponse struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	}

	LoginReq struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	JWTCustomClaims struct {
		Id    int
		Email string
		jwt.StandardClaims
	}
)
