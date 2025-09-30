package response

import "github.com/bagusyanuar/app-pos-be/internal/entity"

type ProductCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToProductCategory(e *entity.ProductCategory) *ProductCategoryResponse {
	return &ProductCategoryResponse{
		ID:   e.ID.String(),
		Name: e.Name,
	}
}

func ToProductCategories(arr []entity.ProductCategory) []ProductCategoryResponse {
	productCategories := make([]ProductCategoryResponse, 0)
	for _, v := range arr {
		productCategory := *ToProductCategory(&v)
		productCategories = append(productCategories, productCategory)
	}
	return productCategories
}
