package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-playground/validator/v10"
	errorCommon "github.com/kmhalpin/todoapp/common/error"
)

func MiddlewareErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs[0]
			if !err.IsType(gin.ErrorTypePrivate) {
				var ves validator.ValidationErrors
				errors.As(err, &ves)
				keys := make(map[string]string)
				for _, ve := range ves {
					keys[ve.Field()] = ve.Tag()
				}
				render.JSON{Data: Error{
					Code:    c.Writer.Status(),
					Message: err.Error(),
					Errors:  keys,
				}}.Render(c.Writer)
				return
			}
			switch err := err.Err.(type) {
			case *errorCommon.ClientError:
				c.JSON(err.StatusCode, Error{
					Code:    err.StatusCode,
					Message: err.Message,
				})
			default:
				if err := errorCommon.TranslateDomainError(err); err != nil {
					c.JSON(err.StatusCode, Error{
						Code:    err.StatusCode,
						Message: err.Message,
					})
				} else {
					c.JSON(http.StatusInternalServerError, Error{
						Code:    http.StatusInternalServerError,
						Message: "Internal server error",
					})
				}
			}
		}
	}
}
