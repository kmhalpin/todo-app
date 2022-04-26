package user

import (
	"context"
)

func (r userUsecase) DeleteUser(ctx context.Context, id string, uid string) (rid string, err error) {
	err = r.GetUserIsRoot(ctx, uid)
	if err != nil {
		return rid, err
	}
	return r.userRepo.DeleteUser(ctx, id)
}
