package models

import "gosql/db"

type Game struct {
	Id 		  int
	Nome      string
	Descricao string
	Preco     float32
	Genero 	  string
}

func BuscaTodosOsJogos() []Game {
	db := db.ConectionSql()
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
	defer db.Close()
	return games
}