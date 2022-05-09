package usecase

import (
	"book-api/model/web"
	"book-api/service"
	"context"

	"gorm.io/gorm"
)

type UserUseCaseImpl struct {
	UserService service.UserService
	DB          *gorm.DB
}

func NewUserUseCase(userService service.UserService, DB *gorm.DB) UserUseCase {
	return &UserUseCaseImpl{
		UserService: userService,
		DB:          DB,
	}
}

func (useCase *UserUseCaseImpl) Create(ctx context.Context, req web.UserCreateReq) (web.UserResponse, error) {
	tx := useCase.DB.Begin()
	user, err := useCase.UserService.Create(ctx, tx, req)
	if err != nil {
		return user, err
	}
	tx.Commit()
	return user, nil
}
