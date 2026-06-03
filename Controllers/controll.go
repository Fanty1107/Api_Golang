package controllers

import(
	"html/template"
	"net/http"
	"gosql/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	AllGames := models.BuscaTodosOsJogos()
	templates.ExecuteTemplate(w, "Index", AllGames)
}