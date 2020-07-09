package daemon

import (
	"fmt"
	"os"
	"os/exec"

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
	label := "com.7geese.tooling"
	plistPath := fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), label)
	logrus.Infof("Stopping %s", label)
	output, err := exec.Command("/bin/launchctl", "stop", plistPath).CombinedOutput()
	logrus.Debugf("Output (launchctl stop): %s", string(output))
	if err != nil {
		logrus.Fatalf("Failed to stop service: %s", err)
	}
}
