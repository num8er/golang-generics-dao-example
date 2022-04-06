package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	println(os.Getenv("DB_CONNECTION_STRING"))
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))

	if err != nil {
		panic(err.Error())
	}

	DB = db
}

func Close() {
	_ = DB.Close()
}
