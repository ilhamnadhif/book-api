package controller

import (
	"book-api/model/web"
	"book-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthControllerImpl struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return &AuthControllerImpl{
		AuthUseCase: authUseCase,
	}
}

func (controller *AuthControllerImpl) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var loginReq web.LoginReq

	if err := c.Bind(&loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(loginReq); err != nil {
		return err
	}

	user, err := controller.AuthUseCase.Login(ctx, loginReq)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, web.WebResponseSuccess(user))

}
