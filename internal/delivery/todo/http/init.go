package http

import (
	"github.com/gin-gonic/gin"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
	jwt "github.com/kmhalpin/todoapp/common/jwt"
	tUCase "github.com/kmhalpin/todoapp/internal/usecase/todo"
)

type HTTPTodoDelivery struct {
	todoUCase tUCase.Usecase
}

func NewHTTPTodoDelivery(router *gin.RouterGroup, todoUCase tUCase.Usecase, j *jwt.JWTManager) HTTPTodoDelivery {
	h := HTTPTodoDelivery{todoUCase: todoUCase}
	router.Use(httpCommon.MiddlewareJWT(j))

	router.GET("", h.getTodo)
	router.GET("/:id", h.getTodoByID)
	router.POST("", h.postTodo)
	router.PUT("/:id", h.putTodoByID)
	router.DELETE("/:id", h.deleteTodoByID)
	return h
}
