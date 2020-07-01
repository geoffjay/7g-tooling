package route

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/handler"

	"github.com/gin-gonic/gin"
)

// Ping routes
func Ping(config *gcontext.Config, r *gin.Engine) error {
	// Simple keep-alive/ping handler
	r.GET(config.VersionedEndpoint("/ping"), handler.Ping())
	return nil
}
