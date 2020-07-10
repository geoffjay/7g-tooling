package transformer

import (
	"github.com/geoffjay/7g-tooling/internal/model"
)

func OneOnOneTemplateGraphQLToVars(template *model.OneOnOneTemplate) map[string]interface{} {
	questions := make([]map[string]interface{}, len(template.Questions))
	for i, question := range template.Questions {
		questions[i] = map[string]interface{}{
			"content":     question.Content,
			"description": question.Description,
		}
	}
	return map[string]interface{}{
		"oneononeTemplate": map[string]interface{}{
			"description": template.Description,
			"name":        template.Name,
			"questions":   questions,
		},
	}
}
