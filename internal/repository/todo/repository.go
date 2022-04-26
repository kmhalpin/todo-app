package todo

import (
	"context"

	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

//go:generate moq -out mock/init.go -pkg mock . Repository

type Repository interface {
	GetTodo(ctx context.Context) (todos []tModel.Todo, err error)
	GetTodoByID(ctx context.Context, id string) (todo tModel.Todo, err error)
	InsertTodo(ctx context.Context, todo tModel.Todo) (id string, err error)
	UpdateTodo(ctx context.Context, todo tModel.Todo) (id string, err error)
	DeleteTodo(ctx context.Context, id string) (rid string, err error)
}
