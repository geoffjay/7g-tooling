package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	key string

	Command = &cobra.Command{
		Use:   "config",
		Short: "Config commands",
		Run:   config,
	}
)

func init() {
	Command.PersistentFlags().StringVar(&key, "key", "version", "configuration setting key")

	Command.AddCommand(getCmd)
	Command.AddCommand(setCmd)
}

func config(cmd *cobra.Command, args []string) {
	logrus.Debug("executing configuration operations")
}
