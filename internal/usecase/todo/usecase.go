package todo

import (
	"context"

	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

type Usecase interface {
	GetTodo(ctx context.Context, uid string) (todos []tModel.Todo, err error)
	GetTodoByID(ctx context.Context, id string, uid string) (todo tModel.Todo, err error)
	InsertTodo(ctx context.Context, todo tModel.Todo, uid string) (id string, err error)
	UpdateTodo(ctx context.Context, todo tModel.Todo, uid string) (id string, err error)
	DeleteTodo(ctx context.Context, id string, uid string) (rid string, err error)
}
