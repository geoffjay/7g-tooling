package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the service",
		Run:   start,
	}
)

func init() {
}

func start(cmd *cobra.Command, args []string) {
	logrus.Debug("start service")
}
