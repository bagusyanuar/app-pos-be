package service

import (
	"context"

	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"github.com/bagusyanuar/app-pos-be/internal/repository"
	"github.com/bagusyanuar/app-pos-be/internal/schema"
	"github.com/shopspring/decimal"
)

type (
	ProductService interface {
		Create(ctx context.Context, schema *schema.ProductSchema) error
	}

	productServiceImpl struct {
		ProductRepository repository.ProductRepository
		Config            *config.AppConfig
	}
)

func NewProductService(
	productRepository repository.ProductRepository,
	config *config.AppConfig,
) ProductService {
	return &productServiceImpl{
		ProductRepository: productRepository,
		Config:            config,
	}
}

// Create implements ProductService.
func (s *productServiceImpl) Create(ctx context.Context, schema *schema.ProductSchema) error {
	productCategoryId := schema.ProductCategoryID
	name := schema.Name
	description := schema.Description
	price := decimal.NewFromFloat(schema.Price)

	e := entity.Product{
		ProductCategoryID: &productCategoryId,
		Name:              name,
		Description:       &description,
		Price:             price,
	}

	_, err := s.ProductRepository.Create(ctx, &e)
	if err != nil {
		return err
	}

	return nil
}
