package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load config:
		//config := service.NewConfig()
		// TODO: send it all back
		c.String(http.StatusOK, "OK")
	}
}

func UpdateConfig() gin.HandlerFunc {
	// Load config:
	//config := service.NewConfig()
	// TODO: do something?
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

func ReadProperty() gin.HandlerFunc {
	// Load config:
	//config := service.NewConfig()
	// TODO: send back one value for a given key
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

func UpdateProperty() gin.HandlerFunc {
	// Load config:
	//config := service.NewConfig()
	// TODO: send a value back for a given key
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
