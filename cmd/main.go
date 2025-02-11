package main

import (
	"log"

	"github.com/Crabocod/golang-test/config"
	"github.com/Crabocod/golang-test/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	a := app.New(cfg)
	if err := a.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
