package user

import (
	passCommon "github.com/kmhalpin/todoapp/common/password"
	uRepo "github.com/kmhalpin/todoapp/internal/repository/user"
)

type userUsecase struct {
	userRepo    uRepo.Repository
	passManager *passCommon.PasswordHashManager
}

func NewUserUsecase(userRepo uRepo.Repository, passManager *passCommon.PasswordHashManager) Usecase {
	return userUsecase{
		userRepo:    userRepo,
		passManager: passManager,
	}
}
