package handler

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	Config *config.AppConfig
}

func NewHomeHandler(cfg *config.AppConfig) *HomeHandler {
	return &HomeHandler{
		Config: cfg,
	}
}

func (c *HomeHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"app_name":    c.Config.Viper.GetString("APP_NAME"),
		"app_version": c.Config.Viper.GetString("APP_VERSION"),
	})
}
