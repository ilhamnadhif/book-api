package controller

import (
	"book-api/helper"
	"book-api/model/web"
	"book-api/usecase"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookControllerImpl struct {
	BookUseCase usecase.BookUseCase
}

func NewBookController(bookUseCase usecase.BookUseCase) BookController {
	return &BookControllerImpl{
		BookUseCase: bookUseCase,
	}
}

func (controller *BookControllerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()
	user := helper.GetCurrentUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("Login Dahulu").Error())
	}
	var bookCreateReq web.BookCreateReq

	if err := c.Bind(&bookCreateReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bookCreateReq.UserID = user.Id
	if err := c.Validate(bookCreateReq); err != nil {
		return err
	}

	book, err := controller.BookUseCase.Create(ctx, bookCreateReq)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(book))

}

func (controller *BookControllerImpl) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	book, err := controller.BookUseCase.FindById(ctx, bookId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(book))
}

func (controller *BookControllerImpl) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	books, err := controller.BookUseCase.FindAll(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(books))
}

func (controller *BookControllerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := helper.GetCurrentUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("Login Dahulu").Error())
	}
	var bookUpdateReq web.BookUpdateReq
	if err := c.Bind(&bookUpdateReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bookUpdateReq.ID = bookId
	if err := c.Validate(bookUpdateReq); err != nil {
		return err
	}
	book, err := controller.BookUseCase.Update(ctx, bookUpdateReq, user.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(book))
}

func (controller *BookControllerImpl) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := helper.GetCurrentUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("Login Dahulu").Error())
	}
	err = controller.BookUseCase.Delete(ctx, bookId, user.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(nil))
}
