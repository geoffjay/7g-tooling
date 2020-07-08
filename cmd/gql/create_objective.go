package gql

import (
	"context"
	"fmt"

	bin "github.com/geoffjay/7g-tooling/data/gql"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ./target/7g gql create-objective --name="foo" --type=1

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
	// --type=1
	createObjectiveCmd.PersistentFlags().IntVarP(&objectiveType, "type", "t", 1, "objective type")
	createObjectiveCmd.PersistentFlags().StringVarP(&name, "name", "n", "Goal", "objective name")
}

func createObjective(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-objective")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	query := bin.GetQuery("create-objective")
	logrus.Debugf("Execute query:\n%s", query)

	objectiveInput := map[string]interface{}{
		"name":          name,
		"objectiveType": objectiveType,
	}

	// type objectiveInputType struct {
	// 	name          string
	// 	objectiveType int
	// }

	// objectiveInput := objectiveInputType{
	// 	name:          name,
	// 	objectiveType: objectiveType,
	// }

	logrus.Debug(objectiveInput)
	client := graphql.NewClient("http://localhost:8000/graphql")
	req := graphql.NewRequest(query)
	req.Var("objective", objectiveInput)
	bearer := fmt.Sprintf("Bearer %s", config.APIKey)
	req.Header.Set("Authorization", bearer)
	ctx := context.Background()
	var res response
	if err := client.Run(ctx, req, &res); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(res)
}
