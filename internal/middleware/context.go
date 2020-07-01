package middleware

import (
	"context"

	"github.com/geoffjay/7g-tooling/internal/util"

	"github.com/gin-gonic/gin"
)

func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), util.ServiceContextKeys.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
