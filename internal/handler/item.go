package handler

import (
	"html/template"
	"net/http"

	"github.com/masaya-nishimura-09/setup-mixer/internal/model"
	"github.com/masaya-nishimura-09/setup-mixer/internal/service"
)

type Keywords struct {
	Title string
}

var templates = template.Must(template.ParseFiles("./templates/item.html"))

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	items := service.SearchItems(q)
	renderTemplate(w, "item", items)
}

func renderTemplate(w http.ResponseWriter, tmpl string, i *model.Items) {
	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
