package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"

	bin "github.com/geoffjay/7g-tooling/data/init"

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
	env := fmt.Sprintf("%s/.config/7g/.env.production", os.Getenv("HOME"))
	data := struct {
		Label     string
		Program   string
		Args      []string
		KeepAlive bool
		RunAtLoad bool
	}{
		Label:     "com.7geese.tooling",
		Program:   "/usr/local/bin/7g",
		Args:      []string{"--env", env, "daemon"},
		KeepAlive: true,
		RunAtLoad: true,
	}

	plistPath := fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), data.Label)
	f, err := os.Create(plistPath)
	if err != nil {
		logrus.Fatalf("Failed to open output file: %s", err)
	}
	tmplName := fmt.Sprintf("%s.plist.tmpl", data.Label)
	logrus.Debugf("Write plist content to %s", plistPath)
	t := template.Must(template.New("launchdConfig").Parse(bin.GetAsset(tmplName)))
	err = t.Execute(f, data)
	if err != nil {
		logrus.Fatalf("Template generation failed: %s", err)
	}

	logrus.Infof("Loading %s", data.Label)
	// We start using load -w on plist file
	output, err := exec.Command("/bin/launchctl", "load", "-w", plistPath).CombinedOutput()
	logrus.Debugf("Output (launchctl load): %s", string(output))
	if err != nil {
		logrus.Fatalf("Failed to load service: %s", err)
	}
}
