package proc

import (
	"github.com/geoffjay/7g-tooling/internal/model"
	"github.com/geoffjay/7g-tooling/pkg/config"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type populateProcessor struct {
	stores map[string]interface{}
}

func NewPopulateProcessor(db *gorm.DB) *populateProcessor {
	return &populateProcessor{
		stores: map[string]interface{}{
			"locations":            model.NewLocationStore(db),
			"departments":          model.NewDepartmentStore(db),
			"users":                model.NewUserStore(db),
			"objectives":           model.NewObjectiveStore(db),
			"one-on-ones":          model.NewOneOnOneStore(db),
			"one-on-one-templates": model.NewOneOnOneTemplateStore(db),
			"recognition-badges":   model.NewRecognitionBadgeStore(db),
			"recognitions":         model.NewRecognitionStore(db),
			"competency-levels":    model.NewLevelStore(db),
			"competencies":         model.NewCompetencyStore(db),
			//"reviews":              model.NewReviewStore(db),
			"roles":          model.NewRoleStore(db),
			"role-templates": model.NewRoleTemplateStore(db),
		},
	}
}

func (p *populateProcessor) ProcessConfig(c *config.Populate) (err error) {
	// TODO: create and return job
	// TODO: create task for each model
	summary := c.Summary()
	for key, _ := range summary {
		switch key {
		case "locations":
			err = p.populateLocations(c.Locations)
		case "departments":
			err = p.populateDepartment(c.Departments)
		case "users":
			err = p.populateUsers(c.Users)
		case "objectives":
			err = p.populateObjectives(c.Objectives)
		case "one-on-ones":
			err = p.populateOneOnOnes(c.OneOnOnes)
		case "one-on-one-templates":
			err = p.populateOneOnOneTemplates(c.OneOnOneTemplates)
		case "recognition-badges":
			err = p.populateRecognitionBadges(c.RecognitionBadges)
		case "recognitions":
			err = p.populateRecognitions(c.Recognitions)
		case "competency-levels":
			err = p.populateLevels(c.CompetencyLevels)
		case "competencies":
			err = p.populateCompetencies(c.Competencies)
		//case "reviews":
		//	err = p.populateReviews(c.Reviews)
		case "roles":
			err = p.populateRoles(c.Roles)
		case "role-templates":
			err = p.populateRoleTemplates(c.RoleTemplates)
		}
	}
	return nil
}

func (p *populateProcessor) populateLocations(locations []*model.Location) error {
	logrus.Debugf("process %d locations", len(locations))
	store := p.stores["locations"].(*model.LocationStore)
	for _, location := range locations {
		if err := store.Save(location); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateDepartment(departments []*model.Department) error {
	logrus.Debugf("process %d departments", len(departments))
	store := p.stores["departments"].(*model.DepartmentStore)
	for _, department := range departments {
		if err := store.Save(department); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateUsers(users []*model.User) error {
	logrus.Debugf("process %d users", len(users))
	store := p.stores["users"].(*model.UserStore)
	for _, user := range users {
		if err := store.Save(user); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateObjectives(objectives []*model.Objective) error {
	logrus.Debugf("process %d objectives", len(objectives))
	store := p.stores["objectives"].(*model.ObjectiveStore)
	for _, objective := range objectives {
		if err := store.Save(objective); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateOneOnOnes(oneOnOnes []*model.OneOnOne) error {
	logrus.Debugf("process %d one one ones", len(oneOnOnes))
	store := p.stores["one-on-ones"].(*model.OneOnOneStore)
	for _, oneOnOne := range oneOnOnes {
		if err := store.Save(oneOnOne); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateOneOnOneTemplates(templates []*model.OneOnOneTemplate) error {
	logrus.Debugf("process %d one on one templates", len(templates))
	store := p.stores["one-on-one-templates"].(*model.OneOnOneTemplateStore)
	for _, template := range templates {
		if err := store.Save(template); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateRecognitionBadges(badges []*model.RecognitionBadge) error {
	logrus.Debugf("process %d badges", len(badges))
	store := p.stores["recognition-badges"].(*model.RecognitionBadgeStore)
	for _, badge := range badges {
		if err := store.Save(badge); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateRecognitions(recognitions []*model.Recognition) error {
	logrus.Debugf("process %d recognitions", len(recognitions))
	store := p.stores["recognitions"].(*model.RecognitionStore)
	for _, recognition := range recognitions {
		if err := store.Save(recognition); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateLevels(levels []*model.Level) error {
	logrus.Debugf("process %d levels", len(levels))
	store := p.stores["competency-levels"].(*model.LevelStore)
	for _, level := range levels {
		if err := store.Save(level); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateCompetencies(competencies []*model.Competency) error {
	logrus.Debugf("process %d competencies", len(competencies))
	store := p.stores["competencies"].(*model.CompetencyStore)
	for _, competency := range competencies {
		if err := store.Save(competency); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateReviews(reviews []*model.Review) error {
	logrus.Debugf("process %d reviews", len(reviews))
	store := p.stores["reviews"].(*model.ReviewStore)
	for _, review := range reviews {
		if err := store.Save(review); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateRoles(roles []*model.Role) error {
	logrus.Debugf("process %d roles", len(roles))
	store := p.stores["roles"].(*model.RoleStore)
	for _, role := range roles {
		if err := store.Save(role); err != nil {
			return err
		}
	}
	return nil
}

func (p *populateProcessor) populateRoleTemplates(templates []*model.RoleTemplate) error {
	logrus.Debugf("process %d role templates", len(templates))
	store := p.stores["role-templates"].(*model.RoleTemplateStore)
	for _, template := range templates {
		if err := store.Save(template); err != nil {
			return nil
		}
	}
	return nil
}
