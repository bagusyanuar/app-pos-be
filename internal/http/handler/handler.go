package handler

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/service"
)

type Handler struct {
	Home            *HomeHandler
	Auth            *AuthHandler
	ProductCategory *ProductCategoryHandler
	Product         *ProductHandler
}

func InitHandler(cfg *config.AppConfig, service *service.Service) *Handler {
	return &Handler{
		Home:            NewHomeHandler(cfg),
		Auth:            NewAuthHandler(service.Auth, cfg),
		ProductCategory: NewProductCategoryHandler(service.ProductCategory, cfg),
		Product:         NewProductHandler(service.Product, cfg),
	}
}
