package user

import (
	"context"

	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

func (r userUsecase) InsertUser(ctx context.Context, user uModel.User, uid string) (id string, err error) {
	err = r.GetUserIsRoot(ctx, uid)
	if err != nil {
		return id, err
	}
	if err := r.userRepo.GetUsernameIsAvailable(ctx, user.Username); err != nil {
		return id, err
	}
	hp, err := r.passManager.HashPassword(user.Password)
	if err != nil {
		return id, err
	}
	user.Password = hp
	return r.userRepo.InsertUser(ctx, user)
}
