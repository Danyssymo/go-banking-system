package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/Danyssymo/go-banking-system/user-service/internal/domain"
	"github.com/Danyssymo/go-banking-system/user-service/internal/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	query := `
		INSERT INTO users (id, email, hashed_password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.pool.Exec(ctx, query,
		user.ID(),
		user.Email(),
		user.HashedPassword(),
		user.CreatedAt(),
		user.UpdatedAt(),
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return ports.ErrEmailAlreadyExists
		}
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	query := `
		SELECT id, email, hashed_password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var (
		id             uuid.UUID
		gotEmail       string
		hashedPassword string
		createdAt      time.Time
		updatedAt      time.Time
	)

	err := r.pool.QueryRow(ctx, query, email).Scan(
		&id,
		&gotEmail,
		&hashedPassword,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ports.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return domain.ReconstituteUser(id, gotEmail, hashedPassword, createdAt, updatedAt), nil
}
