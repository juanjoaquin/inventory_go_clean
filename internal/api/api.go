package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/juanjoaquin/inventory_go_clean/internal/service"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type API struct {
	serv          service.Service     // Con esto manejamos los Users y los Productos
	dataValidator *validator.Validate // Para validar los datos de la request
}

func New(serv service.Service) *API {
	return &API{serv: serv,
		dataValidator: validator.New()}
}

func (a *API) Start(e *echo.Echo, address string) error { // Start va a llamar a nuestro HTTP Framework (Echo)

	/* Registramos las rutas */
	a.RegisterRoutes(e)

	/* Configuramos el CORS */
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:5500"},
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true, // Para que se puedan enviar cookies
		},
	))

	return e.Start(address) // El :8080 es el puerto en el que va a correr nuestro servidor.m
}
