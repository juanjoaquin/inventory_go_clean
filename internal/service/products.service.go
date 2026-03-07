package service

import (
	"context"
	"errors"

	"github.com/juanjoaquin/inventory_go_clean/internal/models"
)

var (
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotFound      = errors.New("product not found")
	ErrInvalidPermissions   = errors.New("user not authorized to add product")
)

var validRolesToAddProduct []int64 = []int64{1, 2}

func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {

	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{} // Declaramos el slice de productos.

	for _, p := range pp {
		products = append(products, models.Product{ // Lo appendeamos al slice de productos.
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			CreatedBy:   p.CreatedBy,
		})
	}

	return products, nil

}

// ME QUEDE EN AGREGAR EL SAVE PRODUCT DEL SERVICE. UWU

func (s *serv) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {

	p, err := s.repo.GetProductByID(ctx, id)

	if err != nil {
		return nil, err
	}

	product := models.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		CreatedBy:   p.CreatedBy,
	}

	return &product, nil

}

func (s *serv) AddProdcut(ctx context.Context, product models.Product, email string) error {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermissions
	}

	return s.repo.SaveProduct(
		ctx,
		product.Name,
		product.Description,
		product.Price,
		u.ID,
	)
}
