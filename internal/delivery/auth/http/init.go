package http

import (
	"github.com/gin-gonic/gin"
	aUCase "github.com/kmhalpin/todoapp/internal/usecase/auth"
)

type HTTPAuthDelivery struct {
	authUCase aUCase.Usecase
}

func NewHTTPAuthDelivery(router *gin.RouterGroup, authUCase aUCase.Usecase) HTTPAuthDelivery {
	h := HTTPAuthDelivery{authUCase: authUCase}

	router.POST("", h.postAuth)
	return h
}
