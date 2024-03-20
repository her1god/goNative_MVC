package config

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("tidak dapat memuat file .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal("DB_PORT DB salah: %v", err)
	}

	konek := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbName + "?parseTime=true"
	db, err := sql.Open("mysql", konek)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// konek := "root:27oktober@/go_produk?parseTime=true"
	// db, err := sql.Open("mysql", konek)
	// if err != nil {
	// 	panic(err)
	// }

	log.Println("Database Connected")
	DB = db

}
