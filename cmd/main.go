// Clean Architecture
// Repository -> Use Case (service) -> Presentation (Rest API / Http Response)

// Settings / Donde correra nuestra aplicacion

// Internal / Todo lo que sea referente a la aplicacion, de forma sensitiva, va dentro de este package

// Sistema de inventario, donde los usuarios tengan diferentes permisos.

// Cree la imagen de docker con:

/*  docker run -d --name mysql-inventory -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:8.0

juan@juan:~/github/inventory_go_clean$ docker logs mysql-inventory */

package main

import (
	"context"
	"fmt"

	"github.com/juanjoaquin/inventory_go_clean/database"
	"github.com/juanjoaquin/inventory_go_clean/internal/api"
	"github.com/juanjoaquin/inventory_go_clean/internal/repository"
	"github.com/juanjoaquin/inventory_go_clean/internal/service"
	"github.com/juanjoaquin/inventory_go_clean/settings"
	"github.com/labstack/echo/v5"
	"go.uber.org/fx"
)

func main() {

	// EL package FX es para la inyeccion de dependencias.
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		), // Pasamos las funciones que devuelven un struct
		fx.Invoke(setLifeCycle),
	)

	app.Run()

}

/* Esta funcion va a ser invocada por el FX y va a ser la que inicie nuestro servidor. Es el ciclo de vida */
/* Iran los hooks para el start y para el stop */

/* Para esto nuevo hay que ejecutar un nuevo docker run */
func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{ // Hook tiene 2 metodos: OnStart y OnStop
		OnStart: func(ctx context.Context) error {

			addres := fmt.Sprintf(":%s", s.Port) // ej. ":8080"
			go a.Start(e, addres)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

}
