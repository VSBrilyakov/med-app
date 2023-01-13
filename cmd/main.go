package main

import (
	"log"

	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/handler"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
	"github.com/mnogohoddovochka/med-app/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(medapp.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server start error: %s", err.Error())
	}
}
