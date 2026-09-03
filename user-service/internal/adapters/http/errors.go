package http

import (
	"errors"
	"net/http"

	"github.com/Danyssymo/go-banking-system/user-service/internal/domain"
	"github.com/Danyssymo/go-banking-system/user-service/internal/ports"
)

func mapError(err error) (int, string) {
	switch {
	case errors.Is(err, ports.ErrEmailAlreadyExists):
		return http.StatusConflict, "Email already exists"
	case errors.Is(err, domain.ErrInvalidEmail):
		return http.StatusBadRequest, "Invalid email"
	case errors.Is(err, domain.ErrEmptyPassword):
		return http.StatusBadRequest, "Empty password"
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}
