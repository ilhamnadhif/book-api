package service

import (
	"book-api/helper"
	"book-api/model/schema"
	"book-api/model/web"
	"book-api/repository"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository  repository.UserRepository
	CacheRepository repository.CacheRepository
}

func NewUserService(userRepository repository.UserRepository, cacheRepository repository.CacheRepository) UserService {
	return &UserServiceImpl{
		UserRepository:  userRepository,
		CacheRepository: cacheRepository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, db *gorm.DB, req web.UserCreateReq) (web.UserResponse, error) {
	hashedPassword, err := helper.BcryptGenerate(req.Password)
	if err != nil {
		return web.UserResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := service.UserRepository.Create(ctx, db, schema.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return web.UserResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, db *gorm.DB, email string) (web.UserResponse, error) {
	user, err := service.UserRepository.FindByEmail(ctx, db, email)
	if err != nil {
		return web.UserResponse{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
