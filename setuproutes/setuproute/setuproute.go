package setuproute

import (
	"go_native/controllers/homecontroller"
	"go_native/controllers/kategoricontroller"
	"go_native/controllers/produkcontroller"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	// Homepage
	r.HandleFunc("/", homecontroller.Welcome)

	// Kategori routes
	kategori := "/kategori"
	r.HandleFunc(kategori, kategoricontroller.Index)
	r.HandleFunc(kategori+"/add", kategoricontroller.Add)
	r.HandleFunc(kategori+"/edit", kategoricontroller.Edit)
	r.HandleFunc(kategori+"/delete", kategoricontroller.Delete)

	// Produk routes
	produk := "/produk"
	r.HandleFunc(produk, produkcontroller.Index)
	r.HandleFunc(produk+"/add", produkcontroller.Add)
	r.HandleFunc(produk+"/detail", produkcontroller.Detail)
	r.HandleFunc(produk+"/edit", produkcontroller.Edit)
	r.HandleFunc(produk+"/delete", produkcontroller.Delete)

	return r
}
