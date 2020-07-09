package transformer

import (
	"github.com/geoffjay/7g-tooling/internal/model"
)

func LocationToGraphQLVars(location *model.Location) map[string]interface{} {
	return map[string]interface{}{
		"name": location.Name,
	}
}
