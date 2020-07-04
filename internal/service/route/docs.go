package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Docs route
func Docs(r *gin.Engine) error {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}
