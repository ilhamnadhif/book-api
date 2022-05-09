package helper

import (
	"book-api/model/web"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetCurrentUser(c echo.Context) *web.JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*web.JWTCustomClaims)
	return claims
}
