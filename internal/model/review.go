package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Review struct {
	gorm.Model
	SgID int
}

type ReviewStore struct {
	db *gorm.DB
}

func NewReviewStore(db *gorm.DB) *ReviewStore {
	return &ReviewStore{
		db: db,
	}
}

func (store *ReviewStore) Save(review *Review) (err error) {
	if err = store.db.Create(review).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *ReviewStore) Flush() (err error) {
	return store.db.Model(&Review{}).Delete(&Review{}).Error
}
