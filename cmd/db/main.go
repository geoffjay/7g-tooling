package db

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "db",
		Short: "Database commands",
		Run:   db,
	}
)

func init() {
	Command.AddCommand(userCmd)
}

func db(cmd *cobra.Command, args []string) {
	logrus.Debug("executing database action")
}
