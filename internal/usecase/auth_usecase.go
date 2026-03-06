package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/ghitufnine/my-go/internal/repository"
	"github.com/ghitufnine/my-go/pkg/jwt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo  repository.UserRepository
	tokenRepo repository.RefreshTokenRepository
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	tokenRepo repository.RefreshTokenRepository,
) *AuthUsecase {
	return &AuthUsecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, email, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	return u.userRepo.Create(ctx, user)
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (string, string, error) {

	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", fmt.Errorf("invalid credentials")
	}

	accessToken, err := jwt.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	err = u.tokenRepo.Store(
		ctx,
		user.ID,
		refreshToken,
		time.Now().Add(7*24*time.Hour),
	)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, token string) error {
	return u.tokenRepo.Delete(ctx, token)
}
