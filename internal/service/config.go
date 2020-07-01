package service

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
)

func NewConfig() *gcontext.Config {
	// Load configuration file
	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("Unable to load configuration: %s\n", err)
	}
	return config
}
