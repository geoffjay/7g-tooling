package daemon

import (
	"fmt"
	"os"
	"os/exec"

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
	label := "com.7geese.tooling"
	plistPath := fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), label)
	logrus.Infof("Unloading %s", label)
	output, err := exec.Command("/bin/launchctl", "unload", plistPath).CombinedOutput()
	logrus.Debugf("Output (launchctl unload): %s", string(output))
	if err != nil {
		logrus.Fatalf("Failed to unload service: %s", err)
	}
}
