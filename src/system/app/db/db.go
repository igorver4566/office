package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	db, err := sql.Open("mysql", "u839560_office:Zxczxc888@tcp(91.217.9.197)/u839560_office")

	if err != nil {
		panic(err.Error())
	}
	return db

}
