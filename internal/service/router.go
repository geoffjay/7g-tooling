package service

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/service/route"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RegisterRoutes register the routes for the server
func RegisterRoutes(config *gcontext.Config, db *gorm.DB, r *gin.Engine) (err error) {
	// 7Geese routes
	if err = route.Network(config, db, r); err != nil {
		return
	}

	// Config routes
	if err = route.Config(config, r); err != nil {
		return
	}

	// Miscellaneous routes
	if err = route.Ping(config, r); err != nil {
		return
	}

	if err = route.Docs(r); err != nil {
		return
	}

	return nil
}
