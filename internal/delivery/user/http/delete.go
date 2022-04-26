package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h HTTPUserDelivery) deleteUserByID(c *gin.Context) {
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

	id, err := h.userUCase.DeleteUser(ctx, id, uid)
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
