package handler

import (
	"errors"

	"github.com/bagusyanuar/app-pos-be/common/exception"
	"github.com/bagusyanuar/app-pos-be/common/util"
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/schema"
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

	data, err := c.ProductCategoryService.FindAll(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get product categories",
		"data":    data,
	})
}

func (c *ProductCategoryHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := c.ProductCategoryService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get product category",
		"data":    data,
	})
}

func (c *ProductCategoryHandler) Create(ctx *fiber.Ctx) error {
	req := new(schema.ProductCategorySchema)
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

	err = c.ProductCategoryService.Create(ctx.UserContext(), &schema)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new product category",
	})
}

func (c *ProductCategoryHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := new(schema.ProductCategorySchema)
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

	err = c.ProductCategoryService.Update(ctx.UserContext(), id, &schema)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully update product category",
	})
}

func (c *ProductCategoryHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.ProductCategoryService.Delete(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully delete product category",
	})
}
