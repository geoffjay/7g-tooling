package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is simple keep-alive/ping handler
// Ping godoc
// @Summary Service check
// @Description Check if the service is alive
// @Success 200 {string} response "OK"
// @Router /ping [get]
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
