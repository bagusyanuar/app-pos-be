package schema

type ProductCategorySchema struct {
	Name string `json:"name" validate:"required"`
}
