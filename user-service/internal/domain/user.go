package domain

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail  = errors.New("invalid email")
	ErrEmptyPassword = errors.New("password hash must not be empty")
)

type User struct {
	id             uuid.UUID
	email          string
	hashedPassword string
	createdAt      time.Time
	updatedAt      time.Time
}

func NewUser(email string, hashedPassword string) (User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return User{}, ErrInvalidEmail
	}

	if hashedPassword == "" {
		return User{}, ErrEmptyPassword
	}

	now := time.Now()

	return User{
		id:             uuid.New(),
		email:          email,
		hashedPassword: hashedPassword,
		createdAt:      now,
		updatedAt:      now,
	}, nil
}

func ReconstituteUser(id uuid.UUID, email string, hashedPassword string, createdAt, updatedAt time.Time) User {
	return User{
		id:             id,
		email:          email,
		hashedPassword: hashedPassword,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Email() string {
	return u.email
}

func (u User) HashedPassword() string {
	return u.hashedPassword
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}
