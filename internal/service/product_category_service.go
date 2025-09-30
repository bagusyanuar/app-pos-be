package service

import (
	"context"

	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"github.com/bagusyanuar/app-pos-be/internal/repository"
	"github.com/bagusyanuar/app-pos-be/internal/response"
	"github.com/bagusyanuar/app-pos-be/internal/schema"
)

type (
	ProductCategoryService interface {
		FindAll(ctx context.Context) (*[]response.ProductCategoryResponse, error)
		FindByID(ctx context.Context, id string) (*response.ProductCategoryResponse, error)
		Create(ctx context.Context, schema *schema.ProductCategorySchema) error
		Update(ctx context.Context, id string, schema *schema.ProductCategorySchema) error
		Delete(ctx context.Context, id string) error
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

// Create implements ProductCategoryService.
func (s *productCategoryServiceImpl) Create(ctx context.Context, schema *schema.ProductCategorySchema) error {
	e := entity.ProductCategory{
		Name: schema.Name,
	}

	_, err := s.ProductCategoryRepository.Create(ctx, &e)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryService.
func (s *productCategoryServiceImpl) FindAll(ctx context.Context) (*[]response.ProductCategoryResponse, error) {
	data, err := s.ProductCategoryRepository.FindAll(ctx)
	if err != nil {
		return &[]response.ProductCategoryResponse{}, err
	}

	res := response.ToProductCategories(data)
	return &res, nil
}

// FindByID implements ProductCategoryService.
func (s *productCategoryServiceImpl) FindByID(ctx context.Context, id string) (*response.ProductCategoryResponse, error) {
	data, err := s.ProductCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := response.ToProductCategory(data)
	return res, nil
}

// Update implements ProductCategoryService.
func (s *productCategoryServiceImpl) Update(ctx context.Context, id string, schema *schema.ProductCategorySchema) error {
	entry := map[string]any{
		"name": schema.Name,
	}

	_, err := s.ProductCategoryRepository.Update(ctx, id, entry)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements ProductCategoryService.
func (s *productCategoryServiceImpl) Delete(ctx context.Context, id string) error {
	err := s.ProductCategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
