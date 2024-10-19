package main

import (
	db "github.com/iput-kernel/foundation-account/internal/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/repository"
	"github.com/iput-kernel/foundation-account/internal/rest/controller"
	"github.com/iput-kernel/foundation-account/internal/rest/router"
	"github.com/iput-kernel/foundation-account/internal/usecase"
	"github.com/iput-kernel/foundation-account/internal/validator"
)

func main() {
	db := db.New()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":4000"))
}
