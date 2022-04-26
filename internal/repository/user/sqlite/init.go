package sqlite

import (
	"database/sql"

	uuidCommon "github.com/kmhalpin/todoapp/common/uuid"
	uRepo "github.com/kmhalpin/todoapp/internal/repository/user"
)

type sqliteUserRepository struct {
	db   *sql.DB
	uuid *uuidCommon.UUIDGenerator
}

func NewSqliteUserRepository(db *sql.DB, uuid *uuidCommon.UUIDGenerator) uRepo.Repository {
	return sqliteUserRepository{
		db:   db,
		uuid: uuid,
	}
}
