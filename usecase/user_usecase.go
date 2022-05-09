package usecase

import (
	"book-api/model/web"
	"context"
)

type UserUseCase interface {
	Create(ctx context.Context, req web.UserCreateReq) (web.UserResponse, error)
}
