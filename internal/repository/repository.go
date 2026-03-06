package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/juanjoaquin/inventory_go_clean/internal/entity"
)

// Repository es la interfaz que engloba las CRUD operations
//
//go:generate mockery --name=Repository --output=repository --inpackage=true
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
