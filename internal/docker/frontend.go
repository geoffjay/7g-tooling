package docker

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func BuildFrontend(cli *client.Client) {
	buildContext, err := createBuildContext("dockerfile.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := buildContext.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	options := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           []string{"7geese-frontend"},
		Dockerfile:     "data/frontend.Dockerfile",
	}
	response, err := cli.ImageBuild(context.Background(), buildContext, options)
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}

func RunFrontend(cli *client.Client) {
	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "7geese-frontend",
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
