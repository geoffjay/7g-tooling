package transformer

import "github.com/geoffjay/7g-tooling/internal/model"

func KeyResultToGraphQLVars(keyResult *model.KeyResult) map[string]interface{} {
	return map[string]interface{}{
		"name":            keyResult.Name,
		"measurementType": keyResult.MeasurementType,
		"weight":          keyResult.Weight,
		"targetValue":     keyResult.TargetValue,
		"startingValue":   keyResult.StartingValue,
	}
}

func ObjectiveToGraphQLVars(objective *model.Objective) map[string]interface{} {
	keyResults := make([]map[string]interface{}, len(objective.KeyResults))
	for i, kr := range objective.KeyResults {
		keyResults[i] = KeyResultToGraphQLVars(kr)
	}
	return map[string]interface{}{
		"objective": map[string]interface{}{
			"name":          objective.Name,
			"objectiveType": objective.GetTypeID(),
			"draft":         objective.Draft,
			"description":   objective.Description,
			"dueDate":       objective.DueDate,
			"startDate":     objective.StartDate,
			//"closed":        objective.Closed,
			"private": objective.Private,
			//"departmentVisibility": objective.DepartmentVisibility,
			//"departments": objective.Departments // objective.GetDepartmentSgIDs()?
			//"labels": objective.Labels,
			//"weights": objective.Weights,
			//"owners": objective.GetOwnerSgIDs(),
			//"stakeholders": objective.GetStakeholderSgIDs(),
			//"followers":  objective.FollowerSgIDs,
			"keyResults": keyResults,
		},
	}
}
