package main

import (
	"log"

	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(medapp.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server start error: %s", err.Error())
	}
}
