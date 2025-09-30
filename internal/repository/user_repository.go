package repository

import (
	"context"
	"errors"

	"github.com/bagusyanuar/app-pos-be/common/exception"
	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindByEmail(ctx context.Context, email string) (*entity.User, error)
	}

	userRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

// FindByEmail implements UserRepository.
func (u *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)

	tx := u.DB.WithContext(ctx)
	if err := tx.Where("email = ?", email).
		First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
