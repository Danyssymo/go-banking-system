package usecase

import (
	"context"
	"errors"

	"github.com/Danyssymo/go-banking-system/user-service/internal/domain"
	"github.com/Danyssymo/go-banking-system/user-service/internal/ports"
)

type RegisterUseCase struct {
	repo   ports.UserRepository
	hasher ports.PasswordHasher
}

func NewRegisterUseCase(repo ports.UserRepository, hasher ports.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, email, password string) (domain.User, error) {
	_, err := uc.repo.FindByEmail(ctx, email)
	if err == nil {
		return domain.User{}, ports.ErrEmailAlreadyExists
	}
	if !errors.Is(err, ports.ErrUserNotFound) {
		return domain.User{}, err
	}

	hashedPassword, err := uc.hasher.Hash(password)
	if err != nil {
		return domain.User{}, err
	}

	user, err := domain.NewUser(email, hashedPassword)
	if err != nil {
		return domain.User{}, err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}
