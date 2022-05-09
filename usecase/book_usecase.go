package usecase

import (
	"book-api/model/web"
	"context"
)

type BookUseCase interface {
	Create(ctx context.Context, req web.BookCreateReq) (web.BookResponse, error)
	FindById(ctx context.Context, bookId int) (web.BookResponse, error)
	FindAll(ctx context.Context) ([]web.BookResponse, error)
	Update(ctx context.Context, req web.BookUpdateReq, currentUserID int) (web.BookResponse, error)
	Delete(ctx context.Context, bookId int, currentUserID int) error
}
