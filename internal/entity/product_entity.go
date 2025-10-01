package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID                uuid.UUID
	ProductCategoryID *uuid.UUID
	Name              string
	Price             decimal.Decimal `gorm:"type:numeric(15,2);default:0;"`
	Description       *string
	Image             *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	ProductCategory   *ProductCategory `gorm:"foreignKey:ProductCategoryID"`
}

func (e *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *Product) TableName() string {
	return "products"
}
