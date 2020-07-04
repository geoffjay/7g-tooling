package cmd

import (
	"log"

	"github.com/geoffjay/7g-tooling/cmd/daemon"
	"github.com/geoffjay/7g-tooling/cmd/deploy"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addCommands()

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("verbose", false)
}

func addCommands() {
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(daemon.Command)
	rootCmd.AddCommand(deploy.Command)
}

func setup() {
	if viper.GetBool("verbose") {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
