package deploy

import (
	"fmt"
	"os"

	"github.com/geoffjay/7g-cli/internal/git"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/geoffjay/7g-cli/internal/docker"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var (
	initializeCmd = &cobra.Command{
		Use:   "initialize",
		Short: "Initialize environment for deployments",
		Run:   initialize,
	}
)

func init() {
}

func initialize(cmd *cobra.Command, args []string) {
	logrus.Debug("initializing...")

	// what should happen?
	// - check for ~/.7g (SG_PATH)
	// - create it if it doesn't exist
	// - load configuration
	// # keep here, stuff above goes to top level?
	// - create $SG_PATH/repository
	// - clone github.com/7Geese/7Geese to $SG_PATH/repository
	// - build docker image adding files from repository
	// - stop and rebuild database containers if --rebuild was given ???
	// - specify image name in stack file? or just use the same one every time
	// - run stack file

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("verbose") {
		docker.DumpContainerList(cli)
	}

	// prepare repository
	// TODO: get path from config? can you use viper.Get after loading file?
	path := fmt.Sprintf("%s/repository", os.Getenv("SG_PATH"))
	git.Clone("https://github.com/7Geese/7Geese", path)
	// TODO: create `deploy run` command that does this:
	// git.Checkout(viper.GetString("branch"))

	docker.BuildBase(cli)
	// TODO: create `deploy run` command that does this:
	docker.BuildBackend(cli)
	docker.BuildFrontend(cli)
	docker.RunBackend(cli)
	docker.RunFrontend(cli)
}
