package main

import (
	"log"

	"github.com/po1nt-1/todo-app-backend/config"
	"github.com/po1nt-1/todo-app-backend/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
