package util

import (
	"net/http"
	"strconv"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	temp, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func ParsingId(idString string) (int, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return id, nil
}
