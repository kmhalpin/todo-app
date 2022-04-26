package sqlite

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/kmhalpin/todoapp/common/error"
	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

func (r sqliteTodoRepository) InsertTodo(ctx context.Context, todo tModel.Todo) (id string, err error) {
	uid, err := r.uuid.Generate()
	if err != nil {
		return id, err
	}
	row := r.db.QueryRowContext(
		ctx,
		"INSERT INTO todos (id, title, note, time, owner) VALUES (?, ?, ?, ?, ?) RETURNING id",
		"todo-"+uid,
		todo.Title,
		todo.Note,
		todo.Time,
		todo.UserID,
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("todo not found")
	}
	return id, err
}
