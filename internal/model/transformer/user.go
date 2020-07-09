package transformer

import (
	"fmt"

	"github.com/geoffjay/7g-tooling/internal/model"
)

func UserToGraphQLVars(user *model.User) map[string]interface{} {
	return map[string]interface{}{
		"userProfile": map[string]interface{}{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"username":  fmt.Sprintf("%s%s", user.FirstName, user.LastName),
		},
	}
}
