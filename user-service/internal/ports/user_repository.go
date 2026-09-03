package ports

import (
	"context"
	"errors"

	"github.com/Danyssymo/go-banking-system/user-service/internal/domain"
)

var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyExists = errors.New("email already exists")

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
}
