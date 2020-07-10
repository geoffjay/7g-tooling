package transformer

import "github.com/geoffjay/7g-tooling/internal/model"

func RoleTemplateToGraphQLVars(roleTemplate *model.RoleTemplate) map[string]interface{} {
	expectations := make([]map[string]interface{}, len(roleTemplate.Responsibilities))
	for i, responsibility := range roleTemplate.Responsibilities {
		expectations[i] = map[string]interface{}{
			"title":          responsibility.Title,
			"description":    responsibility.Description,
			"showAssessment": true,
		}
	}
	return map[string]interface{}{
		"roleTemplate": map[string]interface{}{
			"description":  roleTemplate.Description,
			"title":        roleTemplate.Title,
			"expectations": expectations,
		},
	}
}
