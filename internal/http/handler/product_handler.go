package handler

import (
	"github.com/bagusyanuar/app-pos-be/common/exception"
	"github.com/bagusyanuar/app-pos-be/common/util"
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/schema"
	"github.com/bagusyanuar/app-pos-be/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductService service.ProductService
	Config         *config.AppConfig
}

func NewProductHandler(
	productService service.ProductService,
	config *config.AppConfig,
) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
		Config:         config,
	}
}

func (c *ProductHandler) Create(ctx *fiber.Ctx) error {
	req := new(schema.ProductSchema)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrBodyParser.Error(),
		})
	}

	messages, err := util.Validate(c.Config.Validator, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	schema := *req

	err = c.ProductService.Create(ctx.UserContext(), &schema)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new product",
	})
}
