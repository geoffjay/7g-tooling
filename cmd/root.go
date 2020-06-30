package cmd

import (
	"log"

	"github.com/sirupsen/logrus"

	"github.com/geoffjay/7g-cli/cmd/deploy"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "7g",
		Short: "7Geese CLI",
		Long:  "7Geese developer utilities",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(setup)

	addCommands()

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("verbose", false)
}

func addCommands() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(deploy.Command)
}

func setup() {
	if viper.GetBool("verbose") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debug("load config")
}
