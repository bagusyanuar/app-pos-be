package service

import (
	"context"

	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"github.com/bagusyanuar/app-pos-be/internal/repository"
)

type (
	ProductCategoryService interface {
		FindAll(ctx context.Context) ([]entity.ProductCategory, error)
	}

	productCategoryServiceImpl struct {
		ProductCategoryRepository repository.ProductCategoryRepository
		Config                    *config.AppConfig
	}
)

func NewProductCategoryService(
	productCategoryRepository repository.ProductCategoryRepository,
	config *config.AppConfig,
) ProductCategoryService {
	return &productCategoryServiceImpl{
		ProductCategoryRepository: productCategoryRepository,
		Config:                    config,
	}
}

// FindAll implements ProductCategoryService.
func (s *productCategoryServiceImpl) FindAll(ctx context.Context) ([]entity.ProductCategory, error) {
	data, err := s.ProductCategoryRepository.FindAll(ctx)
	if err != nil {
		return []entity.ProductCategory{}, err
	}
	return data, nil
}
