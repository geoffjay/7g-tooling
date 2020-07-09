package transformer

import (
	"testing"

	"github.com/geoffjay/7g-tooling/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestUserToGraphQLVars(t *testing.T) {
	email := "foo.bar@example.com"
	user := &model.User{
		SgID:              0,
		FirstName:         "Foo",
		LastName:          "Bar",
		Email:             &email,
		Phone:             "555-123-4567",
		ManagerID:         0,
		ManagerName:       "",
		State:             "",
		Title:             "",
		DepartmentsJoined: "",
		LocationsJoined:   "",
		RolesJoined:       "",
		Departments:       nil,
		Locations:         nil,
		Roles:             nil,
	}

	variables := UserToGraphQLVars(user)
	userProfile := make(map[string]interface{})
	userProfile = variables["userProfile"].(map[string]interface{})
	t.Log("It should map the model to GraphQL variables")
	assert.Equal(t, "FooBar", userProfile["username"])
}
