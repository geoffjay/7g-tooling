package service

import (
	"net/http"
	"os"
	"time"

	"github.com/geoffjay/7g-tooling/internal/model"
	"github.com/spf13/viper"

	_ "github.com/geoffjay/7g-tooling/docs"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
)

// @title 7Geese Tooling Service API
// @version 1.0
// @description This is the 7Geese utility server.

// @host localhost:3000
// @BasePath /v1
// @query.collection.format multi

// @x-extension-openapi {"example": "value on a json format"}

// Run launches the server
func Run(config *gcontext.Config) {
	// Set configuration values here
	var (
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	db, err := gcontext.Connect(config)
	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	if viper.GetBool("flush") {
		logrus.Debug("Flushing all database tables")
		stores := map[string]interface{}{
			"locations":            model.NewLocationStore(db),
			"departments":          model.NewDepartmentStore(db),
			"users":                model.NewUserStore(db),
			"objectives":           model.NewObjectiveStore(db),
			"one-on-ones":          model.NewOneOnOneStore(db),
			"one-on-one-templates": model.NewOneOnOneTemplateStore(db),
			"recognition-badges":   model.NewRecognitionBadgeStore(db),
			"recognitions":         model.NewRecognitionStore(db),
			"competency-levels":    model.NewLevelStore(db),
			"competencies":         model.NewCompetencyStore(db),
			"reviews":              model.NewReviewStore(db),
			"roles":                model.NewRoleStore(db),
			"role-templates":       model.NewRoleTemplateStore(db),
		}
		err = stores["locations"].(*model.LocationStore).Flush()
		err = stores["departments"].(*model.DepartmentStore).Flush()
		err = stores["users"].(*model.UserStore).Flush()
		err = stores["objectives"].(*model.ObjectiveStore).Flush()
		err = stores["one-on-ones"].(*model.OneOnOneStore).Flush()
		err = stores["one-on-one-templates"].(*model.OneOnOneTemplateStore).Flush()
		err = stores["recognition-badges"].(*model.RecognitionBadgeStore).Flush()
		err = stores["recognitions"].(*model.RecognitionStore).Flush()
		err = stores["competency-levels"].(*model.LevelStore).Flush()
		err = stores["competencies"].(*model.CompetencyStore).Flush()
		err = stores["roles"].(*model.RoleStore).Flush()
		err = stores["role-templates"].(*model.RoleTemplateStore).Flush()
		if err != nil {
			logrus.Panic(err)
		}
	}

	router := initRouter()

	// Routes and Handlers
	if err := RegisterRoutes(config, db, router); err != nil {
		logrus.Fatal("Failed to register routes:", err)
	}

	// Inform the user where the server is listening
	logrus.Info("Running @ " + config.SchemaVersionedEndpoint(""))

	addr := config.ListenEndpoint()

	// Configure HTTP server
	s := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	// Run the server
	logrus.Infof("Listen for requests on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		logrus.Fatal("server.ListenAndServer:", err)
	}
	logrus.Info("Shutting down")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// Set an 8 MiB memory limit for multipart forms
	router.MaxMultipartMemory = 8 << 20

	// Add middleware
	router.Use(logger.SetLogger())
	router.Use(middleware.Context())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Content-Type",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}
