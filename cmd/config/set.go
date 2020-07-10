package config

import (
	"fmt"
	"strconv"

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

	var value string
	switch key {
	case "host":
		config.Host = sValue
		value = sValue
	case "port":
		config.Port = iValue
		value = string(iValue)
	case "uri-schema":
		config.URISchema = sValue
		value = sValue
	case "version":
		config.Version = sValue
		value = sValue
	case "api-key":
		config.APIKey = sValue
		value = sValue
	case "database/file":
		config.Database.File = sValue
		value = sValue
	case "database/flush":
		config.Database.Flush = bValue
		value = strconv.FormatBool(bValue)
	case "database/auto-migrate":
		config.Database.AutoMigrate = bValue
		value = strconv.FormatBool(bValue)
	case "remote/address":
		config.Remote.Address = sValue
		value = sValue
	case "log/formatter":
		config.Log.Formatter = sValue
		value = sValue
	case "log/level":
		config.Log.Level = sValue
		value = sValue
	}

	fmt.Printf("wrote %s: %s\n", key, value)
	config.Write()
}
