package sqlite

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/kmhalpin/todoapp/common/error"
	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

func (r sqliteTodoRepository) UpdateTodo(ctx context.Context, todo tModel.Todo) (id string, err error) {
	row := r.db.QueryRowContext(
		ctx,
		"UPDATE todos SET title = ?, note = ?, time = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? RETURNING id",
		todo.Title,
		todo.Note,
		todo.Time,
		todo.ID,
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("todo not found")
	}
	return id, err
}
