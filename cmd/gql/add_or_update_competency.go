package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	competencyID          int
	competencyTitle       string
	competencyDescription string
	competencyLevels      []map[string]string

	addOrUpdateCompetencyCmd = &cobra.Command{
		Use:   "add-or-update-competency",
		Short: "Creates a new team or updates an existing one",
		Run:   addOrUpdateCompetency,
	}
)

func init() {
	addOrUpdateCompetencyCmd.PersistentFlags().IntVar(&competencyID, "competencyId", 0, "badge id")
	addOrUpdateCompetencyCmd.PersistentFlags().StringVar(&competencyTitle, "title", "Mickey Mouse Clubhouse", "competency title")
	addOrUpdateCompetencyCmd.PersistentFlags().StringVar(&competencyDescription, "description", "Janitorial Services", "competency description")
}

func addOrUpdateCompetency(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > add-or-update-competency")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	competencyLevels = []map[string]string{{
		"title":       "level one",
		"description": "one",
	}, {
		"title":       "level two",
		"description": "two",
	}}
	expectationTemplate := map[string]interface{}{
		"title":           competencyTitle,
		"description":     competencyDescription,
		"levels":          competencyLevels,
		"expectationType": 1,
	}
	variables := map[string]interface{}{
		"pk":                  competencyID,
		"expectationTemplate": expectationTemplate,
	}
	client.Query("add-or-update-competency", variables, config)
}
