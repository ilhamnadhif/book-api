package usecase

import (
	"book-api/model/web"
	"context"
)

type AuthUseCase interface {
	Login(ctx context.Context, req web.LoginReq) (web.TokenResponse, error)
}
