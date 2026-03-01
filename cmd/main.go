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

	"github.com/jmoiron/sqlx"
	"github.com/juanjoaquin/inventory_go_clean/database"
	"github.com/juanjoaquin/inventory_go_clean/settings"
	"go.uber.org/fx"
)

func main() {

	// EL package FX es para la inyeccion de dependencias.
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		), // Pasamos las funciones que devuelven un struct
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()

}
