package main

import (
	"01-server/internal/app"
	"01-server/internal/config"
	"log"
)

func main() {
	cfg := config.Load()
	if err := app.RunServer(cfg); err != nil {
		log.Fatal(err)
	}
}
