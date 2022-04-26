package user

import (
	"context"
	"errors"

	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

func (r userUsecase) GetUser(ctx context.Context, uid string) (users []uModel.User, err error) {
	err = r.GetUserIsRoot(ctx, uid)
	if err != nil {
		return nil, err
	}
	return r.userRepo.GetUser(ctx)
}

func (r userUsecase) GetUserByID(ctx context.Context, id string, uid string) (user uModel.User, err error) {
	err = r.GetUserIsRoot(ctx, uid)
	if err != nil {
		return user, err
	}
	return r.userRepo.GetUserByID(ctx, id)
}

func (r userUsecase) GetUserIsWriter(ctx context.Context, uid string) (err error) {
	u, err := r.userRepo.GetUserByID(ctx, uid)
	if err != nil {
		return err
	}
	if u.Permission != uModel.Write {
		return errors.New("GET_USER_PERMISSION.NOT_AUTHORIZED")
	}
	return nil
}

func (r userUsecase) GetUserIsReader(ctx context.Context, uid string) (err error) {
	u, err := r.userRepo.GetUserByID(ctx, uid)
	if err != nil {
		return err
	}
	if u.Permission == uModel.Write || u.Permission == uModel.Read {
		return nil
	}
	return errors.New("GET_USER_PERMISSION.NOT_AUTHORIZED")
}

func (r userUsecase) GetUserIsRoot(ctx context.Context, uid string) (err error) {
	u, err := r.userRepo.GetUserByID(ctx, uid)
	if err != nil {
		return err
	}
	if u.Username == uModel.ROOT_USERNAME && u.Permission == uModel.Write {
		return nil
	}
	return errors.New("GET_USER_PERMISSION.NOT_AUTHORIZED")
}
