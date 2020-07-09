package gql

import (
	"fmt"

	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	scopes     []string
	forProfile int

	createStaticAPITokenCmd = &cobra.Command{
		Use:   "create-static-api-token",
		Short: "Create a new static API token",
		Run:   createStaticAPIToken,
	}
)

func init() {
	createStaticAPITokenCmd.PersistentFlags().StringArrayVar(&scopes, "scopes", []string{"all"}, "scopes")
	createStaticAPITokenCmd.PersistentFlags().IntVar(&forProfile, "for-profile", -1, "create for profile ID")
}

func createStaticAPIToken(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-static-api-token")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	if len(scopes) == 0 {
		logrus.Panic("A value for `scopes' is required when creating a new static API token")
	}

	variables := map[string]interface{}{
		"scopes": scopes,
	}
	if forProfile != -1 {
		variables["forProfile"] = forProfile
	}
	res, err := client.Query("create-static-api-token", variables, config)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Print(client.ResponseJSON(res))
}
