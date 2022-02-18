package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDataBase() *sql.DB {
	conectionString := "user=postgres dbname=loja password=Senha@123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conectionString)

	if err != nil {
		panic(err.Error())
	}
	return db
}
