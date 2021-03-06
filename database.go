package gocmn

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(con string) {
	var err error

	db, err = sqlx.Connect("mysql", con)
	if err != nil {
		panic("cannot connect to database")
	}
}

func GetDB() *sqlx.DB {
	return db
}

func CloseDB() {
	db.Close()
}
