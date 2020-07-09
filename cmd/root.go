package cmd

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/geoffjay/7g-tooling/internal/util"

	"github.com/geoffjay/7g-tooling/cmd/gql"

	"github.com/geoffjay/7g-tooling/cmd/daemon"
	"github.com/geoffjay/7g-tooling/cmd/deploy"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	env     string
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

	// Command arguments
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&env, "env", ".env", "environment file")

	// Make available anywhere
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	_ = viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))
	viper.SetDefault("verbose", false)

	// The application environment needs to exist
	util.EnsureAppEnv()

	addCommands()
}

func addCommands() {
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(daemon.Command)
	rootCmd.AddCommand(deploy.Command)
	rootCmd.AddCommand(gql.Command)
}

func setup() {
	if viper.GetBool("verbose") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Load environment
	logrus.Debug("Loading environment from %s", env)
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
