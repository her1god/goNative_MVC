package logic

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	temp, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
