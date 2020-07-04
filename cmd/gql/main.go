package gql

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "gql",
		Short: "Perform GraphQL queries",
		Run:   gql,
	}
)

func init() {
	Command.AddCommand(getObjectiveCmd)
}

func gql(cmd *cobra.Command, args []string) {
	logrus.Debug("executing GraphQL query")
}
