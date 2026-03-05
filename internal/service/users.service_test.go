package service

import (
	"context"
	"os"
	"testing"

	"github.com/juanjoaquin/inventory_go_clean/encryption"
	"github.com/juanjoaquin/inventory_go_clean/internal/entity"
	"github.com/juanjoaquin/inventory_go_clean/internal/repository"
	"github.com/stretchr/testify/mock"
)

// No debemos tener que setear las variables dentro de cada unit test. Hay que hacerlo de forma global.
// Esto se ejecutara al principio de cada unit test.
var repo = &repository.MockRepository{}
var s Service

func TestMain(m *testing.M) {

	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	user := &entity.User{Email: "test@exists.com", Password: encryptedPassword}

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)                       // No deberia retornar un error.
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(user, nil)                    // Deberia retornar un usuario.
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil) // No deberia retornar un error.

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			UserName:      "test",
			Email:         "test@test.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserNameAlreadyExists",
			UserName:      "test",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	// Creamos una nueva referencia de los indices de los test cases.

	ctx := context.Background()
	repo := &repository.MockRepository{}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			s := New(repo)

			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}

func TestLoginUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			UserName:      "test",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidPassword",
			UserName:      "test",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidPassword,
		},
	}

	// Creamos una nueva referencia de los indices de los test cases.

	ctx := context.Background()

	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)                                      // No deberia retornar un error.
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(&entity.User{Email: "test@exists.com"}, nil) // Deberia retornar un usuario.
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)                // No deberia retornar un error.

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			s = New(repo)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}
