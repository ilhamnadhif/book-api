package middleware

import (
	"book-api/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	data, ok := err.(*echo.HTTPError)
	if !ok {
		data = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(data.Code, web.WebResponse{
		Code:   data.Code,
		Status: http.StatusText(data.Code),
		Data:   data.Message,
	})

}
