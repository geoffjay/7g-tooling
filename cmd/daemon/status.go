package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Status of the service",
		Run:   status,
	}
)

func init() {
}

func status(cmd *cobra.Command, args []string) {
	logrus.Debug("status service")
}
