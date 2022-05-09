package usecase

import (
	"book-api/model/web"
	"book-api/service"
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type BookUseCaseImpl struct {
	BookService service.BookService
	DB          *gorm.DB
}

func NewBookUseCase(bookService service.BookService, DB *gorm.DB) BookUseCase {
	return &BookUseCaseImpl{
		BookService: bookService,
		DB:          DB,
	}
}

func (useCase *BookUseCaseImpl) Create(ctx context.Context, req web.BookCreateReq) (web.BookResponse, error) {
	tx := useCase.DB.Begin()
	book, err := useCase.BookService.Create(ctx, tx, req)
	if err != nil {
		return book, err
	}
	tx.Commit()
	return book, nil
}

func (useCase *BookUseCaseImpl) FindById(ctx context.Context, bookId int) (web.BookResponse, error) {
	tx := useCase.DB.Begin()
	book, err := useCase.BookService.FindById(ctx, tx, bookId)
	if err != nil {
		return book, err
	}
	tx.Commit()
	return book, err
}

func (useCase *BookUseCaseImpl) FindAll(ctx context.Context) ([]web.BookResponse, error) {
	tx := useCase.DB.Begin()
	books, err := useCase.BookService.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return books, nil

}

func (useCase *BookUseCaseImpl) Update(ctx context.Context, req web.BookUpdateReq, currentUserID int) (web.BookResponse, error) {
	tx := useCase.DB.Begin()
	findBook, err := useCase.BookService.FindById(ctx, tx, req.ID)
	if err != nil {
		return findBook, err
	}

	if findBook.UserID != currentUserID {
		return findBook, echo.NewHTTPError(http.StatusForbidden, errors.New("Anda tidak dapat mengakses resource ini").Error())
	}

	book, err := useCase.BookService.Update(ctx, tx, req)
	if err != nil {
		return book, err
	}
	tx.Commit()
	return web.BookResponse{
		ID:          book.ID,
		Name:        book.Name,
		Description: book.Description,
		Price:       book.Price,
		CreatedAt:   findBook.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}, nil
}

func (useCase *BookUseCaseImpl) Delete(ctx context.Context, bookId int, currentUserID int) error {
	tx := useCase.DB.Begin()
	book, err := useCase.BookService.FindById(ctx, tx, bookId)
	if err != nil {
		return err
	}
	fmt.Println(currentUserID, book.UserID)
	if book.UserID != currentUserID {
		return echo.NewHTTPError(http.StatusForbidden, errors.New("Anda tidak dapat mengakses resource ini").Error())
	}
	err = useCase.BookService.Delete(ctx, tx, book.ID)
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}
