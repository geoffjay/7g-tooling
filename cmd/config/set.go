package config

import (
	"fmt"

	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	key string

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set configuration property",
		Run:   set,
	}
)

func init() {
	setCmd.PersistentFlags().StringVar(&key, "key", "version", "configuration setting key")
}

func set(cmd *cobra.Command, args []string) {
	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	switch key {
	case "host":
		fmt.Print(config.Host)
	}
}
