package schema

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ProductSchema struct {
	ProductCategoryID uuid.UUID             `form:"product_category_id" validate:"required,uuid4"`
	Name              string                `form:"name" validate:"required"`
	Price             float64               `form:"price" validate:"required,numeric"`
	Description       string                `form:"description"`
	Image             *multipart.FileHeader `fomr:"image"`
}
