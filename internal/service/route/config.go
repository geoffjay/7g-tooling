package route

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/handler"
	"github.com/gin-gonic/gin"
)

// Config route group
func Config(config *gcontext.Config, r *gin.Engine) error {
	group := r.Group(config.VersionedEndpoint("/config"))
	{
		group.GET("", handler.ReadConfig())
		group.POST("", handler.UpdateConfig())
		group.GET("/:key", handler.ReadProperty())
		group.POST("/:key", handler.UpdateProperty())
	}
	return nil
}
