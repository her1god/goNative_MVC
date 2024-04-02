package kategoricontroller

import (
	"go_native/entities"
	"go_native/models/kategorimodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	kategorie := kategorimodel.GetAll()
	data := map[string]any{
		"kategori": kategorie,
	}

	temp, err := template.ParseFiles("views/kategori/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/kategori/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var kategorie entities.Kategori

		kategorie.Name = r.FormValue("name")
		kategorie.CreatedAt = time.Now()
		kategorie.UpdateAt = time.Now()

		if ok := kategorimodel.Create(kategorie); !ok {
			temp, _ := template.ParseFiles("views/kategori/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/kategori", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/kategori/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		categorie := kategorimodel.Detail(id)
		data := map[string]any{
			"kategori": categorie,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var kategorie entities.Kategori
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		kategorie.Name = r.FormValue("name")
		kategorie.UpdateAt = time.Now()

		if ok := kategorimodel.Update(id, kategorie); !ok {
			http.Redirect(w, r, r.Header.Get("referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/kategori", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := kategorimodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/kategori", http.StatusSeeOther)
}
