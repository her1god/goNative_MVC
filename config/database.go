package config

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// ambil dari env go
	err := LoadEnv()
	if err != nil {
		log.Fatalf("Error saat mencoba ambil env: %v", err)
	}

	// get dari method si koneksi
	connectionString := BuildConnectionString()

	// koneksi ke database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error saat koneksi ke database: %v", err)
	}

	// set jumlah max simultan
	db.SetMaxOpenConns(20)

	// Set jumlah max tetap aktif
	db.SetMaxIdleConns(10)

	// verif db
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error verifying database connection: %v", err)
	}

	log.Println("Database Connected")

	DB = db
}

func BuildConnectionString() string {
	return DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + strconv.Itoa(DB_PORT) + ")/" + DB_NAME + "?parseTime=true"
}
