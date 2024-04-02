package produkcontroller

import (
	"go_native/entities"
	"go_native/logic"
	"go_native/models/kategorimodel"
	"go_native/models/produkmodel"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	produks := produkmodel.GetAll()
	data := map[string]any{
		"produk": produks,
	}

	logic.RenderTemplate(w, "views/produk/index.html", data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	produks := produkmodel.Detail(id)
	data := map[string]any{
		"produk": produks,
	}

	logic.RenderTemplate(w, "views/produk/detail.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		kategorie := kategorimodel.GetAll()
		data := map[string]any{
			"kategori": kategorie,
		}

		logic.RenderTemplate(w, "views/produk/create.html", data)
	}

	if r.Method == "POST" {
		var produks entities.Produk

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		produks.Name = r.FormValue("name")
		produks.Kategori.Id = uint(categoryId)
		produks.Stock = int64(stock)
		produks.Description = r.FormValue("description")
		produks.CreatedAt = time.Now()
		produks.UpdatedAt = time.Now()

		if ok := produkmodel.Create(produks); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/produk", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		produks := produkmodel.Detail(id)

		kategorie := kategorimodel.GetAll()
		data := map[string]any{
			"kategori": kategorie,
			"produk":   produks,
		}

		logic.RenderTemplate(w, "views/produk/edit.html", data)
	}

	if r.Method == "POST" {
		var produks entities.Produk

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		produks.Name = r.FormValue("name")
		produks.Kategori.Id = uint(categoryId)
		produks.Stock = int64(stock)
		produks.Description = r.FormValue("description")
		produks.UpdatedAt = time.Now()

		if ok := produkmodel.Update(id, produks); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/produk", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := produkmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/produk", http.StatusSeeOther)
}
