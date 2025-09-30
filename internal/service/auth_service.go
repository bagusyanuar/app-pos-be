package service

import (
	"context"
	"errors"
	"time"

	"github.com/bagusyanuar/app-pos-be/common/exception"
	"github.com/bagusyanuar/app-pos-be/internal/config"
	"github.com/bagusyanuar/app-pos-be/internal/entity"
	"github.com/bagusyanuar/app-pos-be/internal/repository"
	"github.com/bagusyanuar/app-pos-be/internal/schema"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService interface {
		Login(ctx context.Context, schema schema.LoginSchema) (accessToken, refreshToken string, err error)
	}

	authServiceImpl struct {
		UserRepository repository.UserRepository
		Config         *config.AppConfig
	}
)

func NewAuthService(
	userRepository repository.UserRepository,
	config *config.AppConfig,
) AuthService {
	return &authServiceImpl{
		UserRepository: userRepository,
		Config:         config,
	}
}

// Login implements AuthService.
func (a *authServiceImpl) Login(ctx context.Context, schema schema.LoginSchema) (accessToken string, refreshToken string, err error) {
	email := schema.Email
	password := schema.Password

	user, err := a.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", exception.ErrMismatchedpassword
		}
		return "", "", err
	}

	accessToken, err = a.createAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = a.createRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *authServiceImpl) createAccessToken(user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Minute * time.Duration(a.Config.JWT.Expiration))
	claims := jwt.RegisteredClaims{
		Issuer:    a.Config.JWT.Issuer,
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   user.ID.String(),
	}
	accessToken := jwt.NewWithClaims(JWTSignInMethod, claims)
	return accessToken.SignedString([]byte(a.Config.JWT.Secret))
}

func (a *authServiceImpl) createRefreshToken(user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Hour * 24 * time.Duration(a.Config.JWT.ExpirationRefreh))
	claims := jwt.RegisteredClaims{
		Issuer:    a.Config.JWT.Issuer,
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   user.ID.String(),
	}
	refreshToken := jwt.NewWithClaims(JWTSignInMethod, claims)
	return refreshToken.SignedString([]byte(a.Config.JWT.SecretRefresh))
}
