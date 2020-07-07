package route

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Network route group
func Network(config *gcontext.Config, db *gorm.DB, r *gin.Engine) error {
	network := r.Group(config.VersionedEndpoint("/network"))
	{
		network.GET("", handler.ReadNetwork())
		network.POST("/populate", handler.PopulateNetwork(db))
		network.POST("/automate", handler.AutomateNetwork(db))
	}
	return nil
}
