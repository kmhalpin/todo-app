package http

import (
	"github.com/gin-gonic/gin"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
	jwt "github.com/kmhalpin/todoapp/common/jwt"
	uUCase "github.com/kmhalpin/todoapp/internal/usecase/user"
)

type HTTPUserDelivery struct {
	userUCase uUCase.Usecase
}

func NewHTTPUserDelivery(router *gin.RouterGroup, userUCase uUCase.Usecase, j *jwt.JWTManager) HTTPUserDelivery {
	h := HTTPUserDelivery{userUCase: userUCase}
	router.Use(httpCommon.MiddlewareJWT(j))

	router.GET("", h.getUser)
	router.GET("/:id", h.getUserByID)
	router.POST("", h.postUser)
	router.DELETE("/:id", h.deleteUserByID)
	return h
}
