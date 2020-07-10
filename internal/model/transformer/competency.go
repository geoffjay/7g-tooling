package transformer

import (
	"github.com/geoffjay/7g-tooling/internal/model"
)

func CompetencyToGraphQLVars(competency *model.Competency) map[string]interface{} {
	return map[string]interface{}{
		"expectationTemplate": map[string]interface{}{
			"title":           competency.Title,
			"description":     competency.Description,
			"expectationType": competency.GetExpectationType(),
			"levels":          competency.Levels,
		},
	}
}
