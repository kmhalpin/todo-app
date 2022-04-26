package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
)

func (h HTTPUserDelivery) getUser(c *gin.Context) {
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

	user, err := h.userUCase.GetUser(ctx, uid)
	if err != nil {
		c.Error(err)
		return
	}

	data := []httpCommon.User{}
	for _, u := range user {
		data = append(data, h.mapUserModelToResponse(u))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h HTTPUserDelivery) getUserByID(c *gin.Context) {
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

	user, err := h.userUCase.GetUserByID(ctx, id, uid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": h.mapUserModelToResponse(user),
	})
}
