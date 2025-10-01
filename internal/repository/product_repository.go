package repository

import (
	"context"

	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		FindAll(ctx context.Context) ([]entity.Product, error)
		Create(ctx context.Context, e *entity.Product) (*entity.Product, error)
	}

	productRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		DB: db,
	}
}

// FindAll implements ProductRepository.
func (r *productRepositoryImpl) FindAll(ctx context.Context) ([]entity.Product, error) {
	tx := r.DB.WithContext(ctx)

	var data []entity.Product
	if err := tx.Find(&data).Error; err != nil {
		return []entity.Product{}, err
	}

	return data, nil
}

// Create implements ProductRepository.
func (r *productRepositoryImpl) Create(ctx context.Context, e *entity.Product) (*entity.Product, error) {
	tx := r.DB.WithContext(ctx)
	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}
	return e, nil
}
