package gql

import (
	"fmt"

	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	name          string
	objectiveType int

	createObjectiveCmd = &cobra.Command{
		Use:   "create-objective",
		Short: "Create a new objective",
		Run:   createObjective,
	}
)

func init() {
	createObjectiveCmd.PersistentFlags().IntVarP(&objectiveType, "type", "t", 1, "objective type")
	createObjectiveCmd.PersistentFlags().StringVarP(&name, "name", "n", "Goal", "objective name")
}

func createObjective(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-objective")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	objectiveInput := map[string]interface{}{
		"name":          name,
		"objectiveType": objectiveType,
	}
	variables := map[string]interface{}{
		"objective": objectiveInput,
	}
	res, err := client.Query("create-objective", variables, config)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Print(client.ResponseJSON(res))
}
