package api

import "github.com/labstack/echo/v5"

/* Esta funcion va a ser la que defina las rutas de la API */

func (a *API) RegisterRoutes(e *echo.Echo) {

	/* Creamos un grupo de rutas para los usuarios */
	users := e.Group("/users")

	users.POST("/register", a.RegisterUser) /* /users/register */
	users.POST("/login", a.LoginUser)        /* /users/login */
	users.GET("/me", a.Me)                  /* /users/me */

	/* Creamos un grupo de rutas para los productos */
	products := e.Group("/products")

	products.POST("", a.AddProduct) /* /products/ */
}
