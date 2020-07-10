package context

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/geoffjay/7g-tooling/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

func Connect(config *Config) (*gorm.DB, error) {
	var file string
	if config.Database.File != "" {
		file = config.Database.File
		logrus.Debugf("using database file from configuration: %s", file)
	} else {
		// TODO: this should be in home when run as a user, or /etc when run as privileged service
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		file = fmt.Sprintf("%s/.config/7g/data.db", home)
		logrus.Debugf("no database file configured, using default: %s", file)
	}

	db, err := gorm.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if config.Database.AutoMigrate {
		// Migrate the schema
		logrus.Debug("Perform database migrations")
		// 7Geese model related
		err = db.AutoMigrate(&model.Department{}).Error
		err = db.AutoMigrate(&model.Location{}).Error
		err = db.AutoMigrate(&model.User{}).Error
		err = db.AutoMigrate(&model.Competency{}).Error
		err = db.AutoMigrate(&model.Level{}).Error
		err = db.AutoMigrate(&model.Feedback{}).Error
		err = db.AutoMigrate(&model.Objective{}).Error
		err = db.AutoMigrate(&model.KeyResult{}).Error
		err = db.AutoMigrate(&model.OneOnOne{}).Error
		err = db.AutoMigrate(&model.OneOnOneTemplate{}).Error
		err = db.AutoMigrate(&model.Recognition{}).Error
		err = db.AutoMigrate(&model.RecognitionBadge{}).Error
		err = db.AutoMigrate(&model.Review{}).Error
		err = db.AutoMigrate(&model.Role{}).Error
		err = db.AutoMigrate(&model.RoleResponsibility{}).Error
		err = db.AutoMigrate(&model.RoleTemplate{}).Error
		// Job related
		err = db.AutoMigrate(&model.Job{}).Error
		err = db.AutoMigrate(&model.Task{}).Error

		if err != nil {
			logrus.Panic("An error occurred during migration:", err)
		}
	}

	// Flush the database tables if requested
	if viper.GetBool("flush") || config.Database.Flush {
		logrus.Debug("Flushing all database tables")
		stores := map[string]interface{}{
			"locations":            model.NewLocationStore(db),
			"departments":          model.NewDepartmentStore(db),
			"users":                model.NewUserStore(db),
			"objectives":           model.NewObjectiveStore(db),
			"key-results":          model.NewKeyResultStore(db),
			"one-on-ones":          model.NewOneOnOneStore(db),
			"one-on-one-templates": model.NewOneOnOneTemplateStore(db),
			"recognition-badges":   model.NewRecognitionBadgeStore(db),
			"recognitions":         model.NewRecognitionStore(db),
			"competency-levels":    model.NewLevelStore(db),
			"competencies":         model.NewCompetencyStore(db),
			"reviews":              model.NewReviewStore(db),
			"roles":                model.NewRoleStore(db),
			"role-templates":       model.NewRoleTemplateStore(db),
		}
		err = stores["locations"].(*model.LocationStore).Flush()
		err = stores["departments"].(*model.DepartmentStore).Flush()
		err = stores["users"].(*model.UserStore).Flush()
		err = stores["objectives"].(*model.ObjectiveStore).Flush()
		err = stores["key-results"].(*model.KeyResultStore).Flush()
		err = stores["one-on-ones"].(*model.OneOnOneStore).Flush()
		err = stores["one-on-one-templates"].(*model.OneOnOneTemplateStore).Flush()
		err = stores["recognition-badges"].(*model.RecognitionBadgeStore).Flush()
		err = stores["recognitions"].(*model.RecognitionStore).Flush()
		err = stores["competency-levels"].(*model.LevelStore).Flush()
		err = stores["competencies"].(*model.CompetencyStore).Flush()
		err = stores["reviews"].(*model.ReviewStore).Flush()
		err = stores["roles"].(*model.RoleStore).Flush()
		err = stores["role-templates"].(*model.RoleTemplateStore).Flush()
		if err != nil {
			logrus.Panic(err)
		}
	}

	return db, nil
}

func Disconnect(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
