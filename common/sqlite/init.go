package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite() *sql.DB {
	db, err := sql.Open("sqlite3", "configs/db/todo_app.db")

	if err != nil {
		panic(err)
	}

	return db
}
