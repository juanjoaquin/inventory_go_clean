package repository

import (
	"context"

	"github.com/juanjoaquin/inventory_go_clean/internal/entity"
)

const (
	qryInsertProduct = `
	INSERT INTO PRODUCTS (name, description, price, created_by) 
	VALUES (?, ?, ?, ?);
	`

	qryGetProducts = `
	SELECT 
		id, 
		name, 
		description, 
		price, 
		updated_at, 
		updated_by,
		created_by 
	FROM PRODUCTS;
	`

	qryGetProductByID = `
	SELECT 
		id, 
		name, 
		description, 
		price, 
		updated_at, 
		updated_by,
		created_by 
	FROM PRODUCTS 
	WHERE id = ?;
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {

	products := []entity.Product{}

	err := r.db.SelectContext(ctx, &products, qryGetProducts)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repo) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {

	product := &entity.Product{}

	err := r.db.GetContext(ctx, product, qryGetProductByID, id)

	if err != nil {
		return nil, err
	}

	return product, nil

}
