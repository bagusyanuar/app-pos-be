package repository

import (
	"context"

	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"gorm.io/gorm"
)

type (
	ProductCategoryRepository interface {
		FindAll(ctx context.Context) ([]entity.ProductCategory, error)
	}

	productCategoryImpl struct {
		DB *gorm.DB
	}
)

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryImpl{
		DB: db,
	}
}

// FindAll implements ProductCategory.
func (r *productCategoryImpl) FindAll(ctx context.Context) ([]entity.ProductCategory, error) {
	tx := r.DB.WithContext(ctx)

	var data []entity.ProductCategory
	if err := tx.Find(&data).Error; err != nil {
		return []entity.ProductCategory{}, err
	}

	return data, nil
}
