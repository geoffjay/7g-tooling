package config

import (
	"fmt"

	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	sValue string
	iValue int
	bValue bool

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set configuration property",
		Run:   set,
	}
)

func init() {
	Command.PersistentFlags().StringVar(&sValue, "string-value", "", "configuration setting string value")
	Command.PersistentFlags().IntVar(&iValue, "int-value", 0, "configuration setting int value")
	Command.PersistentFlags().BoolVar(&bValue, "bool-value", false, "configuration setting boolean value")
}

func set(cmd *cobra.Command, args []string) {
	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	switch key {
	case "host":
		config.Host = sValue
		//gcontext.WriteConfig(config)
		fmt.Print(config.Host)
	}
}
