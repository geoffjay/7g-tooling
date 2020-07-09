package client

import (
	"context"
	"fmt"

	bin "github.com/geoffjay/7g-tooling/data/gql"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

type response interface{}

func Query(name string, variables map[string]interface{}, config *gcontext.Config) {
	query := bin.GetQuery(name)
	logrus.Debugf("Execute query:\n%s", query)

	client := graphql.NewClient(config.Remote.Address)
	req := graphql.NewRequest(query)
	for key, value := range variables {
		req.Var(key, value)
	}
	bearer := fmt.Sprintf("Bearer %s", config.APIKey)
	req.Header.Set("Authorization", bearer)
	ctx := context.Background()
	var res response
	if err := client.Run(ctx, req, &res); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(res)
}
