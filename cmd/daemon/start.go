package daemon

import (
	"fmt"
	"os"
	"os/exec"

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
	label := "com.7geese.tooling"
	plistPath := fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), label)
	logrus.Infof("Starting %s", label)
	output, err := exec.Command("/bin/launchctl", "start", plistPath).CombinedOutput()
	logrus.Debugf("Output (launchctl start): %s", string(output))
	if err != nil {
		logrus.Fatalf("Failed to start service: %s", err)
	}
}
