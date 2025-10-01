package http

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/http/handler"
	"github.com/bagusyanuar/app-pos-be/internal/http/middleware"
)

func NewRouter(cfg *config.AppConfig, handler *handler.Handler) {
	app := cfg.App

	jwtMiddleware := middleware.VerifyJWT(cfg)

	app.Get("/", handler.Home.Index)
	app.Post("/auth/login", handler.Auth.Login)

	productCategory := app.Group("/product-category", jwtMiddleware)
	productCategory.Get("/", handler.ProductCategory.FindAll)
	productCategory.Post("/", handler.ProductCategory.Create)
	productCategory.Get("/:id", handler.ProductCategory.FindByID)
	productCategory.Put("/:id", handler.ProductCategory.Update)
	productCategory.Delete("/:id", handler.ProductCategory.Delete)

	product := app.Group("/product", jwtMiddleware)
	product.Post("/", handler.Product.Create)
}
