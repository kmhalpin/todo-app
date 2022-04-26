package user

import (
	"context"

	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

//go:generate moq -out mock/init.go -pkg mock . Repository

type Repository interface {
	GetUser(ctx context.Context) (users []uModel.User, err error)
	GetUserByID(ctx context.Context, id string) (user uModel.User, err error)
	GetPasswordByUsername(ctx context.Context, username string) (password string, err error)
	GetIDByUsername(ctx context.Context, username string) (id string, err error)
	GetUsernameIsAvailable(ctx context.Context, username string) (err error)
	InsertUser(ctx context.Context, user uModel.User) (id string, err error)
	DeleteUser(ctx context.Context, id string) (rid string, err error)
}
