package handler

import "github.com/bagusyanuar/app-pos-be/internal/config"

type Handler struct {
	Home *HomeHandler
}

func InitHandler(cfg *config.AppConfig) *Handler {
	return &Handler{
		Home: NewHomeHandler(cfg),
	}
}
