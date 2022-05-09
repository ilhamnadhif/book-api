package controller

import (
	"book-api/model/web"
	"book-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserControllerImpl struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return &UserControllerImpl{
		UserUseCase: userUseCase,
	}
}

func (controller *UserControllerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var userCreateReq web.UserCreateReq

	if err := c.Bind(&userCreateReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userCreateReq); err != nil {
		return err
	}

	user, err := controller.UserUseCase.Create(ctx, userCreateReq)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(user))

}
