// internal/middleware/validator.go
package middleware

import (
	"loan-admin/internal/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := validate.Struct(c.Request); err != nil {
			c.AbortWithStatusJSON(errors.InvalidParameter.Status(), gin.H{
				"error":   errors.InvalidParameter.Message(),
				"details": err.(validator.ValidationErrors),
			})
			return
		}
		c.Next()
	}
}
