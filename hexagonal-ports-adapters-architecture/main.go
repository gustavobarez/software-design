package main

import (
	"context"
	"golang_hexagonal_architecture/adapter/input/controller"
	"golang_hexagonal_architecture/adapter/input/controller/routes"
	"golang_hexagonal_architecture/adapter/output/repository"
	"golang_hexagonal_architecture/application/service"
	"golang_hexagonal_architecture/configuration/logger"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		return
	}

	database, err := pgx.Connect(context.Background(), os.Getenv("DB_URI"))
	if err != nil {
		return
	}

	userController := initDependencies(database)

	r := httprouter.New()
	routes.InitRoutes(r, userController)

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}

func initDependencies(database *pgx.Conn) controller.UserControllerInterface {
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	return controller.NewUserControllerInterface(userService)
}
