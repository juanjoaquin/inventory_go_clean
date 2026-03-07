package service

import (
	"os"
	"testing"

	"github.com/juanjoaquin/inventory_go_clean/encryption"
	"github.com/juanjoaquin/inventory_go_clean/internal/entity"
	"github.com/juanjoaquin/inventory_go_clean/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

/* Este archivo esta creado para usar de forma general el TestMain ya que se puede usar una sola vez por paquete. */
var repo = &repository.MockRepository{}
var s Service

/* Solo puede haber una función TestMain por paquete. Este se debe importar para usar en otros tests*/
func TestMain(m *testing.M) {

	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	user := &entity.User{Email: "test@exists.com", Password: encryptedPassword}

	adminUser := &entity.User{ID: 1, Email: "admin@email.com", Password: encryptedPassword}
	customerUser := &entity.User{ID: 2, Email: "customer@email.com", Password: encryptedPassword}

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)    // No deberia retornar un error.
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(user, nil) // Deberia retornar un usuario.

	repo.On("GetUserByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil) // Get User By Email en Admin.

	repo.On("GetUserByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil) // Get User By Email en Customer.

	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil) // No deberia retornar un error.

	repo.On("GetUserRoles", mock.Anything, int64(1)).Return([]entity.UserRole{{UserID: 1, RoleID: 1}}, nil) // Obtener rol de Admin.

	repo.On("GetUserRoles", mock.Anything, int64(2)).Return([]entity.UserRole{{UserID: 2, RoleID: 3}}, nil) // Obtener rol de Customer

	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil) // No deberia retornar un error.

	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil) // No deberia retornar un error.

	repo.On("SaveProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil) // No deberia retornar un error.

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}
