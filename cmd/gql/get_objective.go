package gql

import (
	"context"
	"fmt"

	bin "github.com/geoffjay/7g-tooling/data/gql"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	//sg := client.NewSevenGeese(config)

	if viper.GetBool("verbose") {
		assets := bin.GetQueryList()
		for asset := range assets {
			logrus.Info(asset)
		}
	}

	query := bin.GetQuery("get-objective")
	logrus.Debugf("Execute query:\n%s", query)

	client := graphql.NewClient("http://localhost:8000/graphql")
	req := graphql.NewRequest(query)
	req.Var("pk", id)
	bearer := fmt.Sprintf("Bearer %s", config.APIKey)
	req.Header.Set("Authorization", bearer)
	ctx := context.Background()
	var res response
	if err := client.Run(ctx, req, &res); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(res)
}
