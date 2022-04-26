package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	errorCommon "github.com/kmhalpin/todoapp/common/error"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func NewJWTManager(AccessTokenKey string) *JWTManager {
	return &JWTManager{
		AccessTokenKey: []byte(AccessTokenKey),
	}
}

func (j JWTManager) GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	return token.SignedString(j.AccessTokenKey)
}

func (j JWTManager) VerifyToken(tokenString string) (string, error) {
	claims := &CustomClaims{}
	if _, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	}); err != nil {
		return "", errorCommon.NewUnauthorizedError("token not valid")
	}

	return claims.ID, nil
}
