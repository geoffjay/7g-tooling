package handler

import (
	"fmt"
	"net/http"

	"github.com/geoffjay/7g-tooling/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

func ReadNetwork() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// PopulateNetwork receives a network definition file in yaml(?) format and initializes the job
//
// Testing with curl:
//
// curl -X POST http://localhost:3000/v1/network/populate \
//  -F "file=@/tmp/test.yaml" \
//  -H "Content-Type: multipart/form-data"
func PopulateNetwork() gin.HandlerFunc {
	return func(c *gin.Context) {
		// the application environment needs to exist
		util.SetupAppEnv()
		file, _ := c.FormFile("file")

		// TODO: this should be in home when run as a user, or /etc when run as privileged service
		home, err := homedir.Dir()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		path := fmt.Sprintf("%s/.config/7g/tmp/network-populate.yml", home)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Debugf("uploaded file %s", file.Filename)
		c.JSON(http.StatusOK, gin.H{"received": file.Filename})
	}
}

// AutomateNetwork receives a network automation profile file in yaml(?) format and initializes the job
//
// Testing with curl:
//
// curl -X POST http://localhost:3000/v1/network/populate \
//  -F "file=@/tmp/test.yaml" \
//  -H "Content-Type: multipart/form-data"
func AutomateNetwork() gin.HandlerFunc {
	return func(c *gin.Context) {
		// the application environment needs to exist
		util.SetupAppEnv()
		file, _ := c.FormFile("file")

		// TODO: this should be in home when run as a user, or /etc when run as privileged service
		home, err := homedir.Dir()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		path := fmt.Sprintf("%s/.config/7g/tmp/network-automate.yml", home)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Debugf("uploaded file %s", file.Filename)
		c.JSON(http.StatusOK, gin.H{"received": file.Filename})
	}
}
