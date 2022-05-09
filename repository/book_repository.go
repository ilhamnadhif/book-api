package repository

import (
	"book-api/model/schema"
	"context"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, db *gorm.DB, book schema.Book) (schema.Book, error)
	FindById(ctx context.Context, db *gorm.DB, bookId int) (schema.Book, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]schema.Book, error)
	Update(ctx context.Context, db *gorm.DB, book schema.Book) (schema.Book, error)
	Delete(ctx context.Context, db *gorm.DB, bookId int) error
}
