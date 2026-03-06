package service

import (
	"context"

	"github.com/juanjoaquin/inventory_go_clean/internal/models"
	"github.com/juanjoaquin/inventory_go_clean/internal/repository"
)

// Service aca sería como el USE CASE. Es la business logic.
//
//go:generate mockery --name=Service --output=service --inpackage=true
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error) // Pedimos el Modelo. Porque aqui necesito retornar el modelo del usuario sin saber el password.
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
