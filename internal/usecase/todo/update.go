package todo

import (
	"context"

	tModel "github.com/kmhalpin/todoapp/internal/model/todo"
)

func (r todoUsecase) UpdateTodo(ctx context.Context, todo tModel.Todo, uid string) (id string, err error) {
	err = r.userUsecase.GetUserIsWriter(ctx, uid)
	if err != nil {
		return id, err
	}
	return r.todoRepo.UpdateTodo(ctx, todo)
}
