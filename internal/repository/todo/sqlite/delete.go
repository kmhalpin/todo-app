package sqlite

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/kmhalpin/todoapp/common/error"
)

func (r sqliteTodoRepository) DeleteTodo(ctx context.Context, id string) (rid string, err error) {
	row := r.db.QueryRowContext(
		ctx,
		"DELETE FROM todos WHERE id = ? RETURNING id",
		id,
	)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, errorCommon.NewNotFoundError("todo not found")
	}
	return rid, err
}
