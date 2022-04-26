package auth

import (
	"context"
)

type Usecase interface {
	LoginUser(ctx context.Context, username string, password string) (accessToken string, err error)
}
