package service

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := gin.Default()

	// Load configuration file
	configDir := filepath.Join(util.RootDir(), "configs")
	config, err := gcontext.LoadConfig(configDir)
	if err != nil {
		t.Fatalf("Unable to load configuration: %s\n", err)
	}

	if err := RegisterRoutes(config, router); err != nil {
		t.Fatal(err)
	}

	// Miscellaneous routes
	t.Run("ping", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/ping")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	// Network routes
	t.Run("network [GET]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/network")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	t.Run("network/populate [POST]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/network/populate")
		req, _ := http.NewRequest("POST", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	t.Run("network/automate [POST]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/network/automate")
		req, _ := http.NewRequest("POST", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	// Config routes
	t.Run("config [GET]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/config")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	t.Run("config [POST]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/config")
		req, _ := http.NewRequest("POST", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	t.Run("config/property [GET]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/config/api-key")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

	t.Run("config/property [POST]", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := config.SchemaVersionedEndpoint("/config/api-key")
		req, _ := http.NewRequest("POST", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})
}
