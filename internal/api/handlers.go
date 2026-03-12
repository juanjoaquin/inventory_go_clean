package api

import (
	"log"
	"net/http"

	"github.com/juanjoaquin/inventory_go_clean/encryption"
	"github.com/juanjoaquin/inventory_go_clean/internal/api/dtos"
	"github.com/juanjoaquin/inventory_go_clean/internal/service"
	"github.com/labstack/echo/v5"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c *echo.Context) error { // c es el contexto de la request

	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	/* Bind es para bindear los parametros de la request a la estructura RegisterUser */
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	/* Usamos el package go validator para validar los parametros de la request: go.get github.com/go-playground/validator/v10 */
	err := a.dataValidator.Struct(params)

	/* Debemos ejecutar nuestro service de register user */
	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password) // Pasamos el param del DTO de Register User (email, name, password..)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	// Devolvemos un error 500 si hay un error distinto an nil
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: err.Error()})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	// O podemos hacer un return de nil simplemente.
	return c.JSON(http.StatusCreated, responseMessage{Message: "User created successfully"})
}

func (a *API) LoginUser(c *echo.Context) error {

	// LLamamos a los parametros de la request
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	/* Validamos los parametros de la request */
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	// Debemos buscar al User en nuestro Service
	user, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	//TODO: Crear el token JWT
	token, err := encryption.SignedLoginToken(user)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	// Lo enviamos mediante cookies
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
	}

	// Pasamos la cookie a la response
	c.SetCookie(cookie)

	//TODO: Retornar el token JWT

	return c.JSON(http.StatusOK, map[string]string{"message": "User logged in successfully"})
}
