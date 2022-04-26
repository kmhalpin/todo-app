package todo

import (
	tRepo "github.com/kmhalpin/todoapp/internal/repository/todo"
	uUCase "github.com/kmhalpin/todoapp/internal/usecase/user"
)

type todoUsecase struct {
	todoRepo    tRepo.Repository
	userUsecase uUCase.Usecase
}

func NewTodoUsecase(todoRepo tRepo.Repository, userUsecase uUCase.Usecase) Usecase {
	return todoUsecase{
		todoRepo:    todoRepo,
		userUsecase: userUsecase,
	}
}
