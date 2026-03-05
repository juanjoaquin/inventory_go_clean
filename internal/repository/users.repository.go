package repository

import (
	"context"

	"github.com/juanjoaquin/inventory_go_clean/internal/entity"
)

const (
	qryInsertUser = `
	
	INSERT INTO USERS (email, name, password) 
	VALUES (?, ?, ?);
	`

	qryGetUserByEmail = `
	SELECT 
		id,
		email, 
		name,
		password
	FROM USERS WHERE email = ?;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	// TODO: Posible error?
	user := &entity.User{}

	// Posible solucion
	//var user *entity.User

	err := r.db.GetContext(ctx, user, qryGetUserByEmail, email)

	if err != nil {
		return nil, err
	}

	return user, err
}
