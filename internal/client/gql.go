package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	bin "github.com/geoffjay/7g-tooling/data/gql"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

type Response interface{}

func Query(name string, variables map[string]interface{}, config *gcontext.Config) (Response, error) {
	query := bin.GetQuery(name)
	logrus.Debugf("Execute query: %s", name)

	apiKey, address := getEnv(config)

	client := graphql.NewClient(address)
	req := graphql.NewRequest(query)
	for key, value := range variables {
		req.Var(key, value)
	}
	bearer := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", bearer)
	var res Response
	if err := client.Run(context.Background(), req, &res); err != nil {
		return res, err
	}

	return res, nil
}

func ResponseJSON(response Response) string {
	bytes, err := json.Marshal(response)
	if err != nil {
		logrus.Error(err)
	}
	return string(bytes)
}

func getEnv(config *gcontext.Config) (string, string) {
	var apiKey string
	var address string
	// Hacky solution to use environment as backup when config isn't available
	if config == nil {
		apiKey = os.Getenv("SG_API_KEY")
		address = os.Getenv("SG_REMOTE_ADDRESS")
		logrus.Debugf("Using secondary configuration from environment")
	} else {
		apiKey = config.APIKey
		address = config.Remote.Address
	}

	return apiKey, address
}

type locationResponse struct {
	Locations struct {
		Edges []struct {
			Node struct {
				Pk int
			}
		}
	}
}

func GetLocationIDByName(name string) (int, error) {
	query := fmt.Sprintf(`query{locations(first:1, search:"%s"){edges{node{pk}}}}`, name)
	apiKey, address := getEnv(nil)

	client := graphql.NewClient(address)
	req := graphql.NewRequest(query)
	bearer := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", bearer)
	var res locationResponse
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return -1, err
	}
	return res.Locations.Edges[0].Node.Pk, nil
}

type departmentResponse struct {
	Teams struct {
		Edges []struct {
			Node struct {
				Pk int
			}
		}
	}
}

func GetDepartmentIDByName(name string) (int, error) {
	query := fmt.Sprintf(`query{teams(first:1, search:"%s"){edges{node{pk}}}}`, name)
	apiKey, address := getEnv(nil)

	client := graphql.NewClient(address)
	req := graphql.NewRequest(query)
	bearer := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", bearer)
	var res departmentResponse
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return -1, err
	}
	return res.Teams.Edges[0].Node.Pk, nil
}

type userResponse struct {
	Profiles struct {
		Edges []struct {
			Node struct {
				Pk int
			}
		}
	}
}

func GetUserIDByEmail(email string) (int, error) {
	query := fmt.Sprintf(`query{profiles(first:1, user_Email:"%s"){edges{node{pk}}}}`, email)
	apiKey, address := getEnv(nil)

	client := graphql.NewClient(address)
	req := graphql.NewRequest(query)
	bearer := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", bearer)
	var res userResponse
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return -1, err
	}
	if len(res.Profiles.Edges) == 0 {
		msg := fmt.Sprintf("No users found with the email %s", email)
		return -1, errors.New(msg)
	}
	return res.Profiles.Edges[0].Node.Pk, nil
}

type objectiveResponse struct {
	Objectives struct {
		Edges []struct {
			Node struct {
				Pk int
			}
		}
	}
}

func GetObjectiveIDByName(name string) (int, error) {
	query := fmt.Sprintf(`query{objectives(last:1, name:"%s"){edges{node{pk}}}}`, name)
	apiKey, address := getEnv(nil)

	client := graphql.NewClient(address)
	req := graphql.NewRequest(query)
	bearer := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", bearer)
	var res objectiveResponse
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return -1, err
	}
	return res.Objectives.Edges[0].Node.Pk, nil
}
