package context

import (
	"fmt"

	"github.com/geoffjay/7g-tooling/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

func Connect() (*gorm.DB, error) {
	// TODO: this should be in home when run as a user, or /etc when run as privileged service
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	file := fmt.Sprintf("%s/.config/7g/data.db", home)
	db, err := gorm.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	logrus.Debug("Perform database migrations")
	// 7Geese model related
	db.AutoMigrate(&model.Department{})
	db.AutoMigrate(&model.Location{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Competency{})
	db.AutoMigrate(&model.Level{})
	db.AutoMigrate(&model.Feedback{})
	db.AutoMigrate(&model.Objective{})
	db.AutoMigrate(&model.OneOnOne{})
	db.AutoMigrate(&model.OneOnOneTemplate{})
	db.AutoMigrate(&model.Recognition{})
	db.AutoMigrate(&model.RecognitionBadge{})
	// Job related
	db.AutoMigrate(&model.Job{})
	db.AutoMigrate(&model.Task{})

	return db, nil
}

func Disconnect(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
