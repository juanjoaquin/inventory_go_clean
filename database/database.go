package database

// Package database es el paquete que se encarga de la conexion a la base de datos.

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/juanjoaquin/inventory_go_clean/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)

	db, err := sqlx.ConnectContext(ctx, "mysql", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db, nil
}
