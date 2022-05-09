package service

import (
	"book-api/model/web"
	"context"

	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx context.Context, db *gorm.DB, req web.UserCreateReq) (web.UserResponse, error)
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (web.UserResponse, error)
}
