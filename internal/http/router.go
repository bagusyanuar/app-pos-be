package http

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/http/handler"
)

func NewRouter(cfg *config.AppConfig, handler *handler.Handler) {
	app := cfg.App

	app.Get("/", handler.Home.Index)
	app.Post("/auth/login", handler.Auth.Login)
}
