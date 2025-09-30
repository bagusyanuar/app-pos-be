package service

import (
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/repository"
)

type Service struct {
	Auth AuthService
}

func InitService(cfg *config.AppConfig, repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.User, cfg),
	}
}
