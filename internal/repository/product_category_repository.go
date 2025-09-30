package repository

import (
	"context"
	"errors"

	"github.com/bagusyanuar/app-pos-be/common/exception"
	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ProductCategoryRepository interface {
		FindAll(ctx context.Context) ([]entity.ProductCategory, error)
		FindByID(ctx context.Context, id string) (*entity.ProductCategory, error)
		Create(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error)
		Update(ctx context.Context, id string, entry map[string]any) (*entity.ProductCategory, error)
		Delete(ctx context.Context, id string) error
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

// FindByID implements ProductCategoryRepository.
func (r *productCategoryImpl) FindByID(ctx context.Context, id string) (*entity.ProductCategory, error) {
	tx := r.DB.WithContext(ctx)

	productCategory, err := r.getCategory(tx, id)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

// Create implements ProductCategoryRepository.
func (r *productCategoryImpl) Create(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error) {
	tx := r.DB.WithContext(ctx)
	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

// Update implements ProductCategoryRepository.
func (r *productCategoryImpl) Update(ctx context.Context, id string, entry map[string]any) (*entity.ProductCategory, error) {
	tx := r.DB.WithContext(ctx)
	productCategory, err := r.getCategory(tx, id)

	if err != nil {
		return nil, err
	}

	if err := tx.Model(productCategory).
		Omit(clause.Associations).
		Updates(&entry).Error; err != nil {
		return nil, err
	}

	return productCategory, nil
}

// Delete implements ProductCategoryRepository.
func (r *productCategoryImpl) Delete(ctx context.Context, id string) error {
	tx := r.DB.WithContext(ctx)
	productCategory, err := r.getCategory(tx, id)

	if err != nil {
		return err
	}

	if err := tx.Delete(productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (r *productCategoryImpl) getCategory(tx *gorm.DB, id string) (*entity.ProductCategory, error) {
	productCategory := new(entity.ProductCategory)

	if err := tx.Where("id = ?", id).
		First(productCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return productCategory, nil
}
