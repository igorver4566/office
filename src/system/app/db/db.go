package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	Db, err := sql.Open("mysql", "u839560_office:Zxczxc888@tcp(91.217.9.197)/u839560_office")

	if err != nil {
		panic(err.Error())
	}
	return Db
}
func Insert(name string, mail string, age string) string {
	d := Init()
	defer d.Close()
	insert, err := d.Query("INSERT INTO Users VALUES ('NULL', '" + name + "', '" + mail + "', '" + age + "')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	return "ok"
}
