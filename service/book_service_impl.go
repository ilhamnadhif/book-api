package service

import (
	"book-api/model/schema"
	"book-api/model/web"
	"book-api/repository"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	"gorm.io/gorm"
)

const (
	findAllBook = "FindAllBook"
)

type BookServiceImpl struct {
	BookRepository  repository.BookRepository
	CacheRepository repository.CacheRepository
}

func NewBookService(bookRepository repository.BookRepository, cacheRepository repository.CacheRepository) BookService {
	return &BookServiceImpl{
		BookRepository:  bookRepository,
		CacheRepository: cacheRepository,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, db *gorm.DB, req web.BookCreateReq) (web.BookResponse, error) {

	book, err := service.BookRepository.Create(ctx, db, schema.Book{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return web.BookResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := service.CacheRepository.Del(ctx, findAllBook); err != nil {
		return web.BookResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return web.BookResponse{
		ID:          book.ID,
		UserID:      book.UserID,
		Name:        book.Name,
		Description: book.Description,
		Price:       book.Price,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}, nil
}

func (service *BookServiceImpl) FindById(ctx context.Context, db *gorm.DB, bookId int) (web.BookResponse, error) {
	book, err := service.BookRepository.FindById(ctx, db, bookId)
	if err != nil {
		return web.BookResponse{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return web.BookResponse{
		ID:          book.ID,
		UserID:      book.UserID,
		Name:        book.Name,
		Description: book.Description,
		Price:       book.Price,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}, nil
}

func (service *BookServiceImpl) FindAll(ctx context.Context, db *gorm.DB) ([]web.BookResponse, error) {
	var booksResp []web.BookResponse
	err := service.CacheRepository.Get(ctx, findAllBook, &booksResp)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if booksResp != nil {
		return booksResp, nil
	}
	booksResp = make([]web.BookResponse, 0)
	books, err := service.BookRepository.FindAll(ctx, db)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for _, book := range books {
		booksResp = append(booksResp, web.BookResponse{
			ID:          book.ID,
			UserID:      book.UserID,
			Name:        book.Name,
			Description: book.Description,
			Price:       book.Price,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		})
	}
	if err := service.CacheRepository.Set(ctx, findAllBook, booksResp); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return booksResp, nil
}

func (service *BookServiceImpl) Update(ctx context.Context, db *gorm.DB, req web.BookUpdateReq) (web.BookResponse, error) {
	book, err := service.BookRepository.Update(ctx, db, schema.Book{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return web.BookResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := service.CacheRepository.Del(ctx, findAllBook); err != nil {
		return web.BookResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return web.BookResponse{
		ID:          book.ID,
		UserID:      book.UserID,
		Name:        book.Name,
		Description: book.Description,
		Price:       book.Price,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}, nil
}

func (service *BookServiceImpl) Delete(ctx context.Context, db *gorm.DB, bookId int) error {
	if err := service.CacheRepository.Del(ctx, findAllBook); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := service.BookRepository.Delete(ctx, db, bookId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
