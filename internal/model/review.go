package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Review struct {
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
