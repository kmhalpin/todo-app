package sqlite

import (
	"database/sql"

	uuidCommon "github.com/kmhalpin/todoapp/common/uuid"
	tRepo "github.com/kmhalpin/todoapp/internal/repository/todo"
)

type sqliteTodoRepository struct {
	db   *sql.DB
	uuid *uuidCommon.UUIDGenerator
}

func NewSqliteTodoRepository(db *sql.DB, uuid *uuidCommon.UUIDGenerator) tRepo.Repository {
	return sqliteTodoRepository{
		db:   db,
		uuid: uuid,
	}
}
