package sqlite

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/kmhalpin/todoapp/common/error"

	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

func (r sqliteTodoRepository) GetTodo(ctx context.Context) (todos []tModel.Todo, err error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM todos ORDER BY time DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos = []tModel.Todo{}
	for rows.Next() {
		var todo tModel.Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Note,
			&todo.Time,
			&todo.UserID,
			&todo.TimeCreated,
			&todo.TimeUpdated,
		); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r sqliteTodoRepository) GetTodoByID(ctx context.Context, id string) (todo tModel.Todo, err error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM todos WHERE id = ?", id)
	err = row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Note,
		&todo.Time,
		&todo.UserID,
		&todo.TimeCreated,
		&todo.TimeUpdated,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return todo, errorCommon.NewNotFoundError("todo not found")
	}
	return todo, err
}
