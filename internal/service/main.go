package service

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Setup database connection
	db, err := gcontext.Connect(config)
	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	// Catch shutdown signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		if err := gcontext.Disconnect(db); err != nil {
			logrus.Error(err)
			os.Exit(-1)
		}
		os.Exit(0)
	}()

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
