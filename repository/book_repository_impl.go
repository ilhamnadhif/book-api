package repository

import (
	"book-api/model/schema"
	"context"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Create(ctx context.Context, db *gorm.DB, book schema.Book) (schema.Book, error) {

	res := db.WithContext(ctx).Create(&book)
	err := res.Error
	if err != nil {
		db.Rollback()
		return book, err
	}
	return book, nil
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, bookId int) (schema.Book, error) {
	var book schema.Book
	res := db.WithContext(ctx).First(&book, "id = ?", bookId)
	err := res.Error
	if err != nil {
		db.Rollback()
		return book, err
	}
	return book, nil
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]schema.Book, error) {
	var books []schema.Book
	res := db.WithContext(ctx).Find(&books)
	err := res.Error
	if err != nil {
		db.Rollback()
		return books, err
	}
	return books, nil
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, db *gorm.DB, book schema.Book) (schema.Book, error) {

	res := db.WithContext(ctx).Updates(&book)
	err := res.Error
	if err != nil {
		db.Rollback()
		return book, err
	}

	return book, nil

}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, bookId int) error {
	res := db.WithContext(ctx).Delete(&schema.Book{}, bookId)
	err := res.Error
	if err != nil {
		db.Rollback()
		return err
	}
	return nil
}
