package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	enableCmd = &cobra.Command{
		Use:   "enable",
		Short: "Enable the service",
		Run:   enable,
	}
)

func init() {
}

func enable(cmd *cobra.Command, args []string) {
	logrus.Debug("enable service")
}
