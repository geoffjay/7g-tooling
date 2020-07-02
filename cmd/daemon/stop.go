package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop the service",
		Run:   stop,
	}
)

func init() {
}

func stop(cmd *cobra.Command, args []string) {
	logrus.Debug("stop service")
}
