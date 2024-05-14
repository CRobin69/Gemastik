package main

import (
	"log"
	"os"

	"github.com/CRobinDev/Gemastik/internal/config"
	"github.com/CRobinDev/Gemastik/internal/handler"
	"github.com/CRobinDev/Gemastik/internal/handler/rest"
	"github.com/CRobinDev/Gemastik/internal/repository"
	"github.com/CRobinDev/Gemastik/internal/service"
	"github.com/CRobinDev/Gemastik/pkg/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("env")
	if err != nil && env == "" {
		log.Fatalf("Failed to load env, err : %v", err)
	}

	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	handler := handler.NewHandler(service)

	middleware := middleware.Init(service)

	rest := rest.NewRest(handler, middleware)

	config.Migrate(db)

	rest.RestRoute()

}
