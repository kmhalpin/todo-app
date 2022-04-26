package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
)

func (h HTTPAuthDelivery) postAuth(c *gin.Context) {
	var body httpCommon.User
	ctx := c.Request.Context()
	if err := c.BindJSON(&body); err != nil {
		return
	}

	accessToken, err := h.authUCase.LoginUser(ctx, body.Username, body.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"accessToken": accessToken,
		},
	})
}
