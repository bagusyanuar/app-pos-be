package service

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/repository"
)

type Service struct {
	Auth            AuthService
	ProductCategory ProductCategoryService
	Product         ProductService
}

func InitService(cfg *config.AppConfig, repo *repository.Repository) *Service {
	return &Service{
		Auth:            NewAuthService(repo.User, cfg),
		ProductCategory: NewProductCategoryService(repo.ProductCategory, cfg),
		Product:         NewProductService(repo.Product, cfg),
	}
}
