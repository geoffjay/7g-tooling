package util

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HasAppEnv() bool {
	// TODO: check if
	//  - $SG_PATH exists
	//  - config file exists
	return false
}

func SetupAppEnv() {
	// TODO:
	//  - create $SG_PATH
	//  - create base config
}

// RootDir returns the root path of the project.
// This only works at this level, if it's moved to a different depth it won't work.
func RootDir() string {
	a, b, c, x := runtime.Caller(0)
	logrus.Debug(a, c, x)
	d := path.Join(path.Dir(b), "..")
	return filepath.Dir(d)
}

// GinContextFromContext can be used to read the Gin context in a request, use with:
//
// gc, err := util.GinContextFromContext(ctx)
// if err != nil {
//   return nil, err
// }
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ServiceContextKeys.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
