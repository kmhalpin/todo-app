package auth

import (
	"context"
)

func (a authUsecase) LoginUser(ctx context.Context, username string, password string) (accessToken string, err error) {
	epass, err := a.userRepo.GetPasswordByUsername(ctx, username)
	if err != nil {
		return accessToken, err
	}

	if err := a.passManager.CheckPasswordHash(password, epass); err != nil {
		return accessToken, err
	}

	id, err := a.userRepo.GetIDByUsername(ctx, username)
	if err != nil {
		return accessToken, err
	}

	return a.jwtManager.GenerateToken(id)
}
