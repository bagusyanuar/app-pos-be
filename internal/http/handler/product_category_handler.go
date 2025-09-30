package handler

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProductCategoryHandler struct {
	ProductCategoryService service.ProductCategoryService
	Config                 *config.AppConfig
}

func NewProductCategoryHandler(
	productCategoryService service.ProductCategoryService,
	config *config.AppConfig,
) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		ProductCategoryService: productCategoryService,
		Config:                 config,
	}
}

func (c *ProductCategoryHandler) FindAll(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get product categories",
	})
}
