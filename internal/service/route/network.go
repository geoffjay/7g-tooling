package route

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/handler"

	"github.com/gin-gonic/gin"
)

// Network route group
func Network(config *gcontext.Config, r *gin.Engine) error {
	network := r.Group(config.VersionedEndpoint("/network"))
	{
		network.GET("", handler.ReadNetwork())
		network.POST("/prepare", handler.PrepareNetwork())
		network.POST("/schedule", handler.ScheduleNetwork())
	}
	return nil
}
