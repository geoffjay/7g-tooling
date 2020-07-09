package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	teamID            int
	teamName          string
	teamProfilesToAdd []int
	teamParent        int

	createOrUpdateTeamCmd = &cobra.Command{
		Use:   "create-or-update-team",
		Short: "Creates a new team or updates an existing one",
		Run:   createOrUpdateTeam,
	}
)

func init() {
	createOrUpdateTeamCmd.PersistentFlags().IntVar(&teamID, "teamID", 2, "team id")
	createOrUpdateTeamCmd.PersistentFlags().StringVar(&teamName, "name", "Engineering", "team name")
	createOrUpdateTeamCmd.PersistentFlags().IntSliceVar(&teamProfilesToAdd, "profilesToAdd", []int{1, 2}, "profiles to add in team")
	createOrUpdateTeamCmd.PersistentFlags().IntVar(&teamParent, "type", 1, "objective type")
}

func createOrUpdateTeam(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-or-update-team")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	variables := map[string]interface{}{
		"name":          teamName,
		"pk":            teamID,
		"profilesToAdd": teamProfilesToAdd,
		"parent":        teamParent,
	}
	client.Query("create-or-update-team", variables, config)
}
