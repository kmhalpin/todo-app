package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
)

func (h HTTPTodoDelivery) getTodo(c *gin.Context) {
	iuid, ok := c.Get("uid")
	if !ok {
		c.Error(ErrorUid)
		return
	}
	uid, ok := iuid.(string)
	if !ok {
		c.Error(ErrorUid)
		return
	}
	ctx := c.Request.Context()

	todo, err := h.todoUCase.GetTodo(ctx, uid)
	if err != nil {
		c.Error(err)
		return
	}

	data := []httpCommon.Todo{}
	for _, t := range todo {
		data = append(data, h.mapTodoModelToResponse(t))
	}

	c.PureJSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h HTTPTodoDelivery) getTodoByID(c *gin.Context) {
	iuid, ok := c.Get("uid")
	if !ok {
		c.Error(ErrorUid)
		return
	}
	uid, ok := iuid.(string)
	if !ok {
		c.Error(ErrorUid)
		return
	}
	ctx := c.Request.Context()
	id := c.Param("id")

	todo, err := h.todoUCase.GetTodoByID(ctx, id, uid)
	if err != nil {
		c.Error(err)
		return
	}

	c.PureJSON(http.StatusOK, gin.H{
		"data": h.mapTodoModelToResponse(todo),
	})
}
