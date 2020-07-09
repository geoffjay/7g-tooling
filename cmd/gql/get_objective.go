package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	id int

	getObjectiveCmd = &cobra.Command{
		Use:   "get-objective",
		Short: "Retrieve an objective by ID",
		Run:   getObjective,
	}
)

func init() {
	Command.PersistentFlags().IntVarP(&id, "id", "", 0, "objective ID")
}

func getObjective(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > get-objective")
	logrus.Debugf("gql > get-objective > %d", id)

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	variables := map[string]interface{}{"pk": id}
	client.Query("get-objective", variables, config)
}
