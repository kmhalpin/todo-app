package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
)

func (h HTTPTodoDelivery) putTodoByID(c *gin.Context) {
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
	var body httpCommon.Todo
	ctx := c.Request.Context()
	if err := c.BindJSON(&body); err != nil {
		return
	}
	body.ID = c.Param("id")

	id, err := h.todoUCase.UpdateTodo(ctx, h.mapTodoBodyToModel(body), uid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}
