package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "config",
		Short: "Config commands",
		Run:   config,
	}
)

func init() {
	Command.AddCommand(getCmd)
	Command.AddCommand(setCmd)
}

func config(cmd *cobra.Command, args []string) {
	logrus.Debug("executing configuration operations")
}
