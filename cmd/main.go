package main

import (
	"log"

	medapp "github.com/mnogohoddovochka/med-app"
)

func main() {
	srv := new(medapp.Server)

	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Server start error: %s", err.Error())
	}
}
