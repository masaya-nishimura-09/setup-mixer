package handler

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")
	t.Execute(w, nil)
}
