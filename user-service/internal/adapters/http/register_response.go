package http

import (
	"time"

	"github.com/google/uuid"

	"github.com/Danyssymo/go-banking-system/user-service/internal/domain"
)

type RegisterResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newRegisterResponse(user domain.User) RegisterResponse {
	return RegisterResponse{
		ID:        user.ID(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
	}
}
