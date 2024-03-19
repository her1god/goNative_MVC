package entities

import "time"

type Produk struct {
	Id uint
	Name string
	Kategori Kategori
	Stock int64
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}