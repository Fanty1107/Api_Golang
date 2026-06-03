package db

import 
(
_ "github.com/lib/pq"
	"database/sql"
)

func ConectionSql() *sql.DB {
	conection := "user=postgres dbname=fanty_shop password=061298edu host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}