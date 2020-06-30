package docker

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/jhoonb/archivex"
)

func createBuildContext(tmpFile string) (buildContext *os.File, err error) {
	tar := new(archivex.TarFile)
	if err = tar.Create(tmpFile); err != nil {
		return nil, err
	}
	if err := tar.AddAll("data", true); err != nil {
		return nil, err
	}
	if err := tar.Close(); err != nil {
		return nil, err
	}

	buildContext, err = os.Open(tmpFile)
	if err != nil {
		return nil, err
	}

	return
}

func DumpContainerList(cli *client.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	logrus.Info("\ncontainers:")
	for _, container := range containers {
		logrus.Infof("%s: %s\n", container.Image, container.Status)
	}
}
