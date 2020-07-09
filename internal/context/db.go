package context

import (
	"fmt"

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
		err = db.AutoMigrate(&model.OneOnOne{}).Error
		err = db.AutoMigrate(&model.OneOnOneTemplate{}).Error
		err = db.AutoMigrate(&model.Recognition{}).Error
		err = db.AutoMigrate(&model.RecognitionBadge{}).Error
		err = db.AutoMigrate(&model.Review{}).Error
		err = db.AutoMigrate(&model.Role{}).Error
		err = db.AutoMigrate(&model.RoleTemplate{}).Error
		// Job related
		err = db.AutoMigrate(&model.Job{}).Error
		err = db.AutoMigrate(&model.Task{}).Error

		if err != nil {
			logrus.Panic("An error occurred during migration:", err)
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
