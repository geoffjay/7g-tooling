package transformer

import (
	"github.com/geoffjay/7g-tooling/internal/model"
)

func CompetencyToGraphQLVars(competency *model.Competency) map[string]interface{} {
	levels := make([]map[string]interface{}, len(competency.Levels))
	for i, level := range competency.Levels {
		levels[i] = map[string]interface{}{
			"title":       level.Title,
			"description": level.Description,
		}
	}
	return map[string]interface{}{
		"expectationTemplate": map[string]interface{}{
			"description":     competency.Description,
			"title":           competency.Title,
			"expectationType": competency.GetExpectationType(),
			"levels":          levels,
		},
	}
}
