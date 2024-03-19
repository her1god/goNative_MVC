package main

import (
	"go_native/config"
	"go_native/controllers/homecontroller"
	"go_native/controllers/kategoricontroller"
	"go_native/controllers/produkcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. kategori
	http.HandleFunc("/kategori", kategoricontroller.Index)
	http.HandleFunc("/kategori/add", kategoricontroller.Add)
	http.HandleFunc("/kategori/edit", kategoricontroller.Edit)
	http.HandleFunc("/kategori/delete", kategoricontroller.Delete)

	// 3. produk
	http.HandleFunc("/produk", produkcontroller.Index)
	http.HandleFunc("/produk/add", produkcontroller.Add)
	http.HandleFunc("/produk/detail", produkcontroller.Detail)
	http.HandleFunc("/produk/edit", produkcontroller.Edit)
	http.HandleFunc("/produk/delete", produkcontroller.Delete)


	log.Println("Server Running At Port 8080")
	http.ListenAndServe(":8080", nil)
	
}