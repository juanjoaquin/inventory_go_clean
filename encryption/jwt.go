package encryption

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/juanjoaquin/inventory_go_clean/internal/models"
)

func SignedLoginToken(user *models.User) (string, error) {

	// Necesitamos obtener la info del User y luego la devolveremos sifrada en el JWT

	// El metodo 256 es viable si el servidor que creo el token, es el encargado de que va a validar el token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
	})

	// El signing string es el string que se va a usar para signar el token. Lo genera como string basicamente
	jwtString, err := token.SignedString([]byte(key /* Es la KEY del archivo de encrypton.go */))

	if err != nil {
		return "", err
	}

	return jwtString, nil

}
