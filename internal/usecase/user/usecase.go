package user

import (
	"context"

	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

type Usecase interface {
	GetUserIsRoot(ctx context.Context, uid string) (err error)
	GetUserIsWriter(ctx context.Context, uid string) (err error)
	GetUserIsReader(ctx context.Context, uid string) (err error)
	GetUser(ctx context.Context, uid string) (users []uModel.User, err error)
	GetUserByID(ctx context.Context, id string, uid string) (user uModel.User, err error)
	InsertUser(ctx context.Context, user uModel.User, uid string) (id string, err error)
	DeleteUser(ctx context.Context, id string, uid string) (rid string, err error)
}
