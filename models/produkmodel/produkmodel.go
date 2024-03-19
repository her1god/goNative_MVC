package produkmodel

import (
	"go_native/config"
	"go_native/entities"
)

func GetAll() []entities.Produk {
	rows, err := config.DB.Query(`
	SELECT
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var produks []entities.Produk

	for rows.Next(){
		var produk entities.Produk
		err := rows.Scan(
			&produk.Id,
			&produk.Name,
			&produk.Kategori.Name,
			&produk.Stock,
			&produk.Description,
			&produk.CreatedAt,
			&produk.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		produks = append(produks, produk)
	}

	return produks
}

func Create(produks entities.Produk) bool {
	result, err := config.DB.Exec(`
	INSERT INTO products(
		name, category_id, stock, description, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?)`,
		produks.Name,
		produks.Kategori.Id,
		produks.Stock,
		produks.Description,
		produks.CreatedAt,
		produks.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}


func Detail(id int) entities.Produk {
	row := config.DB.QueryRow(`
	SELECT
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	WHERE products.id = ?
	`, id)

	var produk entities.Produk
	err := row.Scan(
		&produk.Id,
		&produk.Name,
		&produk.Kategori.Name,
		&produk.Stock,
		&produk.Description,
		&produk.CreatedAt,
		&produk.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return produk
}

func Update(id int, produks entities.Produk) bool {
	query, err := config.DB.Exec(`
	UPDATE products SET
		name = ?,
		category_id = ?,
		stock = ?,
		description = ?,
		updated_at = ?
	WHERE id = ?
	`,
	produks.Name,
	produks.Kategori.Id,
	produks.Stock,
	produks.Description,
	produks.UpdatedAt,
	id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}