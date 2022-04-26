package auth

import (
	jwt "github.com/kmhalpin/todoapp/common/jwt"
	passCommon "github.com/kmhalpin/todoapp/common/password"
	uRepo "github.com/kmhalpin/todoapp/internal/repository/user"
)

type authUsecase struct {
	userRepo    uRepo.Repository
	passManager *passCommon.PasswordHashManager
	jwtManager  *jwt.JWTManager
}

func NewAuthUsecase(userRepo uRepo.Repository, passManager *passCommon.PasswordHashManager, j *jwt.JWTManager) Usecase {
	return authUsecase{
		userRepo:    userRepo,
		passManager: passManager,
		jwtManager:  j,
	}
}
