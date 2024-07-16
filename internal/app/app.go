package app

import (
	"gourlshortener/internal/handler"
	"net/http"
)

type App struct {
    Router  http.Handler
}

func NewApp() *App {
	mux := http.NewServeMux()
	handler.SetupRoutes(mux)
	logMux := handler.LoggingMiddleware(mux);

	return &App{
		Router:  logMux,
	}
}