// Clean Architecture
// Repository -> Use Case (service) -> Presentation (Rest API / Http Response)

// Settings / Donde correra nuestra aplicacion

// Internal / Todo lo que sea referente a la aplicacion, de forma sensitiva, va dentro de este package

package main

import (
	"log"

	"github.com/juanjoaquin/inventory_go_clean/settings"
	"go.uber.org/fx"
)

func main() {

	// EL package FX es para la inyeccion de dependencias.
	app := fx.New(
		fx.Provide(
			settings.New,
		), // Pasamos las funciones que devuelven un struct
		fx.Invoke(
			func(s *settings.Settings) {
				log.Println(s)
			}, // Ejecutamos un comando antes de que la aplicacion se ejecute
		),
	)

	app.Run()

}
