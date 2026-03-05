package service

import (
	"context"
	"errors"

	"github.com/juanjoaquin/inventory_go_clean/encryption"
	"github.com/juanjoaquin/inventory_go_clean/internal/models"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	user, _ := s.repo.GetUserByEmail(ctx, email)

	if user != nil {
		return ErrUserAlreadyExists
	}

	// Hash the password
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	password = encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {

	user, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	// TODO: Decrypt the password
	bb, err := encryption.FromBase64(user.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)

	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidPassword
	}

	/* 	if user.Password != password {
		return nil, ErrInvalidPassword
	} */

	return &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil

}
