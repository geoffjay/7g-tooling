package service

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/service/route"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes register the routes for the server
func RegisterRoutes(config *gcontext.Config, r *gin.Engine) (err error) {
	// Miscellaneous routes
	if err = route.Ping(config, r); err != nil {
		return err
	}

	return err
}
