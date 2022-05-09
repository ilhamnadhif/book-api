package service

import (
	"book-api/model/web"
	"context"

	"gorm.io/gorm"
)

type BookService interface {
	Create(ctx context.Context, db *gorm.DB, req web.BookCreateReq) (web.BookResponse, error)
	FindById(ctx context.Context, db *gorm.DB, bookId int) (web.BookResponse, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]web.BookResponse, error)
	Update(ctx context.Context, db *gorm.DB, req web.BookUpdateReq) (web.BookResponse, error)
	Delete(ctx context.Context, db *gorm.DB, bookId int) error
}
