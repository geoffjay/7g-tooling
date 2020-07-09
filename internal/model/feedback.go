package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Feedback struct {
	gorm.Model
	SgID int
}

type FeedbackStore struct {
	db *gorm.DB
}

func NewFeedbackStore(db *gorm.DB) *FeedbackStore {
	return &FeedbackStore{
		db: db,
	}
}

func (store *FeedbackStore) Save(feedback *Feedback) (err error) {
	if err = store.db.Create(feedback).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *FeedbackStore) Flush() (err error) {
	return store.db.Model(&Feedback{}).Delete(&Feedback{}).Error
}
