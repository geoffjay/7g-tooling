package client

import (
	"context"
	"fmt"
	"os"

	bin "github.com/geoffjay/7g-tooling/data/gql"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

type response interface{}

func Query(name string, variables map[string]interface{}, config *gcontext.Config) {
	query := bin.GetQuery(name)
	logrus.Debugf("Execute query:\n%s", query)

	var apiKey string
	var address string
	// Hacky solution to use environment as backup when config isn't available
	if config == nil {
		apiKey = os.Getenv("SG_API_KEY")
		address = os.Getenv("SG_REMOTE_ADDRESS")
		logrus.Info(address)
		logrus.Debugf("Using secondary configuration from environment")
	} else {
		apiKey = config.APIKey
		address = config.Remote.Address
	}

	client := graphql.NewClient(address)
	req := graphql.NewRequest(query)
	for key, value := range variables {
		req.Var(key, value)
	}
	bearer := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", bearer)
	ctx := context.Background()
	var res response
	if err := client.Run(ctx, req, &res); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(res)
}
