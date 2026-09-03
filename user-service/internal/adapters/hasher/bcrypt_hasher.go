package hasher

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/Danyssymo/go-banking-system/user-service/internal/ports"
)

type BcryptHasher struct{}

var _ ports.PasswordHasher = (*BcryptHasher)(nil)

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
