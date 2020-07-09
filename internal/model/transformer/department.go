package transformer

import (
	"github.com/geoffjay/7g-tooling/internal/model"
)

func DepartmentToGraphQLVars(department *model.Department) map[string]interface{} {
	return map[string]interface{}{
		"name": department.Name,
	}
}
