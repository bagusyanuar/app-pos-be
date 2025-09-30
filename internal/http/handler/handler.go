package handler

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/service"
)

type Handler struct {
	Home *HomeHandler
	Auth *AuthHandler
}

func InitHandler(cfg *config.AppConfig, service *service.Service) *Handler {
	return &Handler{
		Home: NewHomeHandler(cfg),
		Auth: NewAuthHandler(service.Auth, cfg),
	}
}
