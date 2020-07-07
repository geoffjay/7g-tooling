package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Feedback struct {
	BaseModelSeq
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
