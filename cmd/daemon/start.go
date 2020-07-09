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
	//label := "com.7geese.tooling"
	//wait := 30
	//
	//if wait > 0 {
	//	status, waitErr := s.WaitForStatus(wait, 100*time.Millisecond)
	//	if waitErr != nil {
	//		return waitErr
	//	}
	//	if status == nil {
	//		return fmt.Errorf("%s is not running", label)
	//	}
	//	logrus.Debug("Service status: %#v", status)
	//}
}
