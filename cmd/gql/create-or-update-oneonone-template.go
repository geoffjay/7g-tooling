package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	templateID          int
	templateName        string
	templateDescription string
	templateQuestions   []string

	createOrUpdateOneononeTemplateCmd = &cobra.Command{
		Use:   "create-or-update-oneonone-template",
		Short: "Creates a new oneonone template or updates an existing one",
		Run:   createOrUpdateOneononeTemplate,
	}
)

func init() {
	createOrUpdateOneononeTemplateCmd.Flags().IntVar(&templateID, "templateID", 0, "template id")
	createOrUpdateOneononeTemplateCmd.PersistentFlags().StringVar(&templateName, "name", "template", "template name")
	createOrUpdateOneononeTemplateCmd.PersistentFlags().StringVar(&templateDescription, "description", "sss blah", "team description")
	// createOrUpdateOneononeTemplateCmd.PersistentFlags().StringSliceVar(&questions, "questions", []string{"content: q, description: d"}, "questions to add into template")
}

func createOrUpdateOneononeTemplate(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-or-update-oneonone-template")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	// FIXME: read in questions somehow
	templateQuestions := []map[string]string{{
		"content":     "qsaasq",
		"description": "dq",
	}, {
		"content":     "qssq",
		"description": "dq",
	}}

	template := map[string]interface{}{
		"name":        templateName,
		"description": templateDescription,
		"questions":   templateQuestions,
	}

	variables := map[string]interface{}{
		"pk":               templateID,
		"oneononeTemplate": template,
	}
	client.Query("create-or-update-oneonone-template", variables, config)
}
