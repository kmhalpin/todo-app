package todo

import (
	"context"
)

func (r todoUsecase) DeleteTodo(ctx context.Context, id string, uid string) (rid string, err error) {
	err = r.userUsecase.GetUserIsWriter(ctx, uid)
	if err != nil {
		return rid, err
	}
	return r.todoRepo.DeleteTodo(ctx, id)
}
