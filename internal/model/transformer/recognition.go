package transformer

import (
	"github.com/geoffjay/7g-tooling/internal/model"
)

func BadgeToGraphQLVars(badge *model.RecognitionBadge) map[string]interface{} {
	return map[string]interface{}{
		"badge": map[string]interface{}{
			"name":        badge.Name,
			"description": badge.Description,
		},
	}
}
