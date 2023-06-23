package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/po1nt-1/todo-app-backend/config"
	v1 "github.com/po1nt-1/todo-app-backend/internal/controller/http/v1"
	"github.com/po1nt-1/todo-app-backend/internal/httpserver"
	"github.com/po1nt-1/todo-app-backend/internal/usecase"
)

func Run(cfg *config.Config) {
	taskUseCase := usecase.New()

	handler := gin.New()
	v1.NewRouter(handler, taskUseCase)
	httpServer := httpserver.New(
		handler,
		httpserver.Addr(cfg.HTTP.Host, cfg.HTTP.Port),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	select {
	case s := <-interrupt:
		log.Printf("%v Run signal: %v", cfg.App.Name, s)
	case err := <-httpServer.Notify():
		log.Fatalf("%v Run httpServer.Notify: %v", cfg.App.Name, err)
	}

	err := httpServer.Shutdown()
	if err != nil {
		log.Fatalf("%v Run httpServer.Shutdown: %v", cfg.App.Name, err)
	}
}
