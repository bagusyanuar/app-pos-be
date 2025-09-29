package bootstrap

import (
	"fmt"

	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/http"
	"github.com/bagusyanuar/app-pos-be/internal/http/handler"
)

func initialize() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	logger := config.NewLogger(viper)
	defer logger.Sync()
	cfgDB := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(cfgDB)
	cfgJWT := config.NewJWTManager(viper)
	validator := config.NewValidator()

	redisClient := config.NewRedisClient(viper)

	return &config.AppConfig{
		App:       app,
		Viper:     viper,
		DB:        db,
		Logger:    logger,
		JWT:       cfgJWT,
		Validator: validator,
		Redis:     redisClient,
	}
}

func Start() {
	cfg := initialize()

	// start dependency injection
	diHandler := handler.InitHandler(cfg)

	http.NewRouter(cfg, diHandler)
	envPort := cfg.Viper.GetString("APP_PORT")
	port := fmt.Sprintf(":%s", envPort)
	server := cfg.App
	fmt.Println("Fiber server running on", port)
	if err := server.Listen(port); err != nil {
		panic(err)
	}
}
