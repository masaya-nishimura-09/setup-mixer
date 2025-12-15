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

func Nav() []model.ItemNav {
    itemNav := []model.ItemNav{
        {Link: "#desk", Name: "デスク"},
        {Link: "#chair", Name: "イス"},
        {Link: "#keyboard", Name: "キーボード"},
        {Link: "#mouse", Name: "マウス"},
        {Link: "#mouse-pad", Name: "マウスパッド"},
        {Link: "#monitor", Name: "モニター"},
        {Link: "#monitor-arm", Name: "モニターアーム"},
        {Link: "#desk-lamp", Name: "デスクライト"},
        {Link: "#power-strip", Name: "電源タップ"},
        {Link: "#power-strip-case", Name: "ケーブル収納"},
        {Link: "#speaker", Name: "スピーカー"},
        {Link: "#camera", Name: "カメラ"},
        {Link: "#headphone", Name: "ヘッドセット"},
    }
    return itemNav
}


func ItemHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    itemPage := model.ItemPage{
        Nav: Nav(),
        Categories: service.SearchItems(q),
    }
    renderTemplate(w, "item", &itemPage)
}

func renderTemplate(w http.ResponseWriter, tmpl string, i *model.ItemPage) {
    err := templates.ExecuteTemplate(w, tmpl+".html", i)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
