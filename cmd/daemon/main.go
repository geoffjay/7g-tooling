package daemon

import (
	"github.com/geoffjay/7g-tooling/internal/service"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	flush bool

	Command = &cobra.Command{
		Use:   "daemon",
		Short: "Control the 7Geese tooling service",
		Run:   daemon,
	}
)

func init() {
	Command.PersistentFlags().BoolVarP(&flush, "flush", "f", false, "flush all database tables")
	_ = viper.BindPFlag("flush", Command.PersistentFlags().Lookup("flush"))
	viper.SetDefault("flush", false)

	Command.AddCommand(startCmd)
	Command.AddCommand(stopCmd)
	Command.AddCommand(statusCmd)
	Command.AddCommand(enableCmd)
	Command.AddCommand(disableCmd)
}

func daemon(cmd *cobra.Command, args []string) {
	config := service.NewConfig()
	logrus.Debugf("running daemon service on %s\n", config.SchemaVersionedEndpoint(""))

	// TODO: add a --background and use a go routine
	service.Run(config)
}
