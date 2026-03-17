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
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error

	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id int64) (*models.Product, error)
	AddProduct(ctx context.Context, product models.Product, userEmail string) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
