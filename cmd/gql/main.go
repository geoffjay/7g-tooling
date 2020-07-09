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

type response interface{}

func init() {
	// Create and update mutations
	Command.AddCommand(createObjectiveCmd)
	Command.AddCommand(createStaticAPITokenCmd)
	Command.AddCommand(createOrUpdateTeamCmd)
	Command.AddCommand(createUserProfileCmd)

	// Read queries
	Command.AddCommand(getObjectiveCmd)
}

func gql(cmd *cobra.Command, args []string) {
	logrus.Debug("executing GraphQL query")
}
