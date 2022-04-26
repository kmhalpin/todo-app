package sqlite

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/kmhalpin/todoapp/common/error"
	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

func (r sqliteUserRepository) InsertUser(ctx context.Context, user uModel.User) (id string, err error) {
	uid, err := r.uuid.Generate()
	if err != nil {
		return id, err
	}
	row := r.db.QueryRowContext(
		ctx,
		"INSERT INTO users (id, permission, username, password) VALUES (?, ?, ?, ?) RETURNING id",
		"user-"+uid,
		user.GetUserPermissionString(),
		user.Username,
		user.Password,
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}
