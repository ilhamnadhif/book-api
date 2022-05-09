package middleware

import (
	"github.com/go-playground/locales/id_ID"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	id := id_ID.New()
	uni := ut.New(id, id)

	trans, _ := uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(cv.Validator, trans)

	err := cv.Validator.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorResp []string
		for _, e := range errs {
			errorResp = append(errorResp, strings.ToLower(e.Translate(trans)))
		}
		return echo.NewHTTPError(http.StatusBadRequest, strings.Join(errorResp, ", "))
	}
	return nil
}
