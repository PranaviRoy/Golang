package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //mysql driver for go
)

//Conn connects to Database
func Conn() (db *sql.DB) {
	
	connString := "user:password@tcp(localhost)/booksgrpc"
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}
	return db
}
