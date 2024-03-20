package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	konek := "root:27oktober@/go_produk?parseTime=true"
	db, err := sql.Open("mysql", konek)
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db

}
