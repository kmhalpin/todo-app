package http

import (
	"github.com/gin-gonic/gin"
	errorCommon "github.com/kmhalpin/todoapp/common/error"
	jwt "github.com/kmhalpin/todoapp/common/jwt"
)

func MiddlewareJWT(j *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= BEARER {
			c.Error(errorCommon.NewInvariantError("authorization header not valid"))
			c.Abort()
			return
		}
		tokenString := authHeader[BEARER:]

		id, err := j.VerifyToken(tokenString)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("uid", id)
		c.Next()
	}
}
