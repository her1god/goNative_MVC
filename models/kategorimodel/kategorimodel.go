package kategorimodel

import (
	"go_native/config"
	"go_native/entities"
)

func GetAll() []entities.Kategori {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var kategorie []entities.Kategori

	for rows.Next() {
		var kategori entities.Kategori
		if err := rows.Scan(&kategori.Id, &kategori.Name, &kategori.CreatedAt, &kategori.UpdateAt); err != nil {
			panic(err)
		}

		kategorie = append(kategorie, kategori)
	}

	return kategorie
}

func Create(kategorie entities.Kategori) bool {
	result, err := config.DB.Exec(`
	INSERT INTO categories (name, created_at, updated_at)
	VALUE (?, ?, ?)`,
		kategorie.Name, kategorie.CreatedAt, kategorie.UpdateAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Kategori {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var kategorie entities.Kategori
	err := row.Scan(&kategorie.Id, &kategorie.Name)
	if err != nil {
		panic(err)
	}

	return kategorie
}

func Update(id int, kategori entities.Kategori) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, kategori.Name, kategori.UpdateAt, id)
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
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
