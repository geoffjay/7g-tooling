package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	badgeID          int
	badgeName        string
	badgeDescription string

	createOrUpdateBadgeCmd = &cobra.Command{
		Use:   "create-or-update-badge",
		Short: "Creates a new team or updates an existing one",
		Run:   createOrUpdateBadge,
	}
)

func init() {
	createOrUpdateBadgeCmd.PersistentFlags().IntVar(&badgeID, "badgeID", 0, "badge id")
	createOrUpdateBadgeCmd.PersistentFlags().StringVar(&badgeName, "name", "WOW", "badge name")
	createOrUpdateBadgeCmd.PersistentFlags().StringVar(&badgeDescription, "description", "amazing job", "badge description")
}

func createOrUpdateBadge(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-or-update-badge")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	badge := map[string]string{
		"name":        badgeName,
		"description": badgeDescription,
	}
	variables := map[string]interface{}{
		"pk":    badgeID,
		"badge": badge,
	}
	client.Query("create-or-update-badge", variables, config)
}
