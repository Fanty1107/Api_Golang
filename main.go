package main

import (
	"database/sql"
	"net/http"
	"text/template"
	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Game struct {
	Id 		  int
	Nome      string
	Descricao string
	Preco     float32
	Genero    string
}

func ConectionSql() *sql.DB {
	conection := "user=... dbname=... password=... host=... sslmode=..."
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := ConectionSql()
	defer db.Close()
	selectAllGames, err := db.Query("select * from games")
	if err != nil {
		panic(err.Error())
	}

	g := Game{}
	games := []Game{}

	for selectAllGames.Next(){
		var id int
		var Nome, Descricao, Genero string
		var Preco float32

		err = selectAllGames.Scan(&id, &Nome, &Descricao, &Preco, &Genero)
		if err != nil{
			panic(err.Error())
		}
		g.Nome = Nome
		g.Genero = Genero
		g.Descricao = Descricao
		g.Preco = Preco

		games = append(games, g)
	}

	templates.ExecuteTemplate(w, "Index", games)
}
