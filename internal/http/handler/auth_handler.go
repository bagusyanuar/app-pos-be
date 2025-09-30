package handler

import (
	"errors"

	"github.com/bagusyanuar/app-pos-be/common/exception"
	"github.com/bagusyanuar/app-pos-be/common/util"
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/response"
	"github.com/bagusyanuar/app-pos-be/internal/schema"
	"github.com/bagusyanuar/app-pos-be/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService service.AuthService
	Config      *config.AppConfig
}

func NewAuthHandler(
	authService service.AuthService,
	config *config.AppConfig,
) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Config:      config,
	}
}

func (c *AuthHandler) Login(ctx *fiber.Ctx) error {
	request := new(schema.LoginSchema)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrBodyParser.Error(),
		})
	}

	messages, err := util.Validate(c.Config.Validator, request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	schema := *request

	accessToken, refreshToken, err := c.AuthService.Login(ctx.UserContext(), schema)
	if err != nil {
		if errors.Is(err, exception.ErrUserNotFound) {
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

	res := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully login",
		"data":    res,
	})
}
