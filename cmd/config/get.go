package config

import (
	"fmt"

	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get configuration property",
		Run:   get,
	}
)

func init() {
}

func get(cmd *cobra.Command, args []string) {
	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	switch key {
	case "host":
		fmt.Println(config.Host)
	case "port":
		fmt.Println(config.Port)
	case "uri-schema":
		fmt.Println(config.URISchema)
	case "version":
		fmt.Println(config.Version)
	case "api-key":
		fmt.Println(config.APIKey)
	case "database/file":
		fmt.Println(config.Database.File)
	case "database/flush":
		fmt.Println(config.Database.Flush)
	case "database/auto-migrate":
		fmt.Println(config.Database.AutoMigrate)
	case "remote/address":
		fmt.Println(config.Remote.Address)
	case "log/formatter":
		fmt.Println(config.Log.Formatter)
	case "log/level":
		fmt.Println(config.Log.Level)
	}
}
