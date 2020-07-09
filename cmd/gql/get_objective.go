package gql

import (
	"fmt"

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
	Command.PersistentFlags().IntVar(&id, "id", 0, "objective ID")
}

func getObjective(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > get-objective")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	variables := map[string]interface{}{"pk": id}
	res, err := client.Query("get-objective", variables, config)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Print(client.ResponseJSON(res))
}
