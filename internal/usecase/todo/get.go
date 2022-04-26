package todo

import (
	"context"

	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

func (r todoUsecase) GetTodo(ctx context.Context, uid string) (todos []tModel.Todo, err error) {
	err = r.userUsecase.GetUserIsReader(ctx, uid)
	if err != nil {
		return nil, err
	}
	return r.todoRepo.GetTodo(ctx)
}

func (r todoUsecase) GetTodoByID(ctx context.Context, id string, uid string) (todo tModel.Todo, err error) {
	err = r.userUsecase.GetUserIsReader(ctx, uid)
	if err != nil {
		return todo, err
	}
	return r.todoRepo.GetTodoByID(ctx, id)
}
