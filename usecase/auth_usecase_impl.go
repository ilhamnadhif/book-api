package usecase

import (
	"book-api/config"
	"book-api/helper"
	"book-api/model/web"
	"book-api/service"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type AuthUseCaseImpl struct {
	UserService service.UserService
	DB          *gorm.DB
	JWTConfig   config.JWTConfig
}

func NewAuthUseCase(userService service.UserService, DB *gorm.DB, jwtConfig config.JWTConfig) AuthUseCase {
	return &AuthUseCaseImpl{
		UserService: userService,
		DB:          DB,
		JWTConfig:   jwtConfig,
	}
}

func (useCase *AuthUseCaseImpl) Login(ctx context.Context, req web.LoginReq) (web.TokenResponse, error) {
	tx := useCase.DB.Begin()
	user, err := useCase.UserService.FindByEmail(ctx, tx, req.Email)
	if err != nil {
		return web.TokenResponse{}, err
	}
	tx.Commit()
	if !helper.BcryptValidate(user.Password, req.Password) {
		return web.TokenResponse{}, echo.NewHTTPError(http.StatusUnauthorized, errors.New("password not match").Error())
	}

	token, err := helper.GenerateJWTToken(useCase.JWTConfig, web.JWTCustomClaims{
		Id:    user.ID,
		Email: user.Email,
	})
	if err != nil {
		return web.TokenResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return web.TokenResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}, nil
}
