package entity

import "time"

type Product struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float32   `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedBy   int64     `db:"created_by"`
}
