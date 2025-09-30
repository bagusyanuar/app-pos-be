package repository

import "github.com/bagusyanuar/app-pos-be/internal/config"

type Repository struct {
	User            UserRepository
	ProductCategory ProductCategoryRepository
}

func InitRepository(cfg *config.AppConfig) *Repository {
	return &Repository{
		User:            NewUserRepository(cfg.DB),
		ProductCategory: NewProductCategoryRepository(cfg.DB),
	}
}
