package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	disableCmd = &cobra.Command{
		Use:   "disable",
		Short: "Disable the service",
		Run:   disable,
	}
)

func init() {
}

func disable(cmd *cobra.Command, args []string) {
	logrus.Debug("disable service")
}
