package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	roleTemplateTitle       string
	roleTemplateDescription string
	roleExpectations        []map[string]string

	addRoleTemplateCmd = &cobra.Command{
		Use:   "add-role-template",
		Short: "Creates a new team or updates an existing one",
		Run:   addRoleTemplate,
	}
)

func init() {
	addRoleTemplateCmd.PersistentFlags().StringVar(&roleTemplateTitle, "title", "Cook", "role title")
	addRoleTemplateCmd.PersistentFlags().StringVar(&roleTemplateDescription, "description", "At the Mickey Mouse Clubhouse", "role description")
}

func addRoleTemplate(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > add-role-template")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	roleExpectations = []map[string]string{{
		"title":       "Make bananas",
		"description": "in kitchen",
	}, {
		"title":       "Make mangos",
		"description": "in kitchen",
	}}
	roleTemplate := map[string]interface{}{
		"title":        roleTemplateTitle,
		"description":  roleTemplateDescription,
		"expectations": roleExpectations,
	}
	variables := map[string]interface{}{
		"roleTemplate": roleTemplate,
	}
	client.Query("add-role-template", variables, config)
}
