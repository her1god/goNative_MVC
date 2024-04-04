package kategoricontroller

import (
	"go_native/entities"
	"go_native/models/kategorimodel"
	"go_native/util"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	kategorie := kategorimodel.GetAll()
	data := map[string]any{
		"kategori": kategorie,
	}
	util.RenderTemplate(w, "views/kategori/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.RenderTemplate(w, "views/kategori/create.html", nil)
	}

	if r.Method == "POST" {
		var kategorie entities.Kategori

		kategorie.Name = r.FormValue("name")
		kategorie.CreatedAt = time.Now()
		kategorie.UpdateAt = time.Now()

		if ok := kategorimodel.Create(kategorie); !ok {
			util.RenderTemplate(w, "views/kategori/create.html", nil)
		}

		http.Redirect(w, r, "/kategori", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := util.ParsingId(idString)
		if err != nil {
			panic(err)
		}

		categorie := kategorimodel.Detail(id)
		data := map[string]any{
			"kategori": categorie,
		}
		util.RenderTemplate(w, "views/kategori/edit.html", data)
	}

	if r.Method == "POST" {
		var kategorie entities.Kategori
		idString := r.FormValue("id")
		id, err := util.ParsingId(idString)
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
	id, err := util.ParsingId(idString)
	if err != nil {
		panic(err)
	}

	if err := kategorimodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/kategori", http.StatusSeeOther)
}
