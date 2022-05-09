package main

import (
	"book-api/app"
	"book-api/config"
	"book-api/controller"
	custommiddleware "book-api/middleware"
	"book-api/model/web"
	"book-api/repository"
	"book-api/service"
	"book-api/usecase"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	validate := validator.New()
	initConfig := config.Init(validate)

	dbGorm := app.InitGorm(initConfig.Database)
	redis := app.InitRedis(initConfig.Redis)

	redisRepository := repository.NewCacheRepository(redis, time.Minute*time.Duration(initConfig.Redis.TTL))
	bookRepository := repository.NewBookRepository()
	userRepository := repository.NewUserRepository()

	bookService := service.NewBookService(bookRepository, redisRepository)
	userService := service.NewUserService(userRepository, redisRepository)

	bookUseCase := usecase.NewBookUseCase(bookService, dbGorm)
	userUseCase := usecase.NewUserUseCase(userService, dbGorm)
	authUseCase := usecase.NewAuthUseCase(userService, dbGorm, initConfig.JWT)

	bookController := controller.NewBookController(bookUseCase)
	userController := controller.NewUserController(userUseCase)
	authController := controller.NewAuthController(authUseCase)

	jwtConfig := middleware.JWTConfig{
		Claims:     &web.JWTCustomClaims{},
		SigningKey: []byte(initConfig.JWT.SecretKey),
	}

	e := echo.New()
	e.HTTPErrorHandler = custommiddleware.CustomHTTPErrorHandler
	e.Validator = &custommiddleware.CustomValidator{Validator: validate}
	e.Use(middleware.Recover())

	e.Static("/docs", "swagger")

	e.POST("/login", authController.Login)

	e.POST("/users", userController.Create)

	booksRouter := e.Group("/books")
	booksRouter.GET("/:bookId", bookController.FindById)
	booksRouter.GET("", bookController.FindAll)

	booksRouter.Use(middleware.JWTWithConfig(jwtConfig))

	booksRouter.POST("", bookController.Create)
	booksRouter.PUT("/:bookId", bookController.Update)
	booksRouter.DELETE("/:bookId", bookController.Delete)

	e.Logger.Fatal(e.Start(":4000"))
}
