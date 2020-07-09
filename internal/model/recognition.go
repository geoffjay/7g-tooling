package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Recognition struct {
	BaseModelSeq
	Recipient User
	Provider  User
	Badge     RecognitionBadge
	Message   *string `gorm:"size:1024"`
}

type RecognitionStore struct {
	db *gorm.DB
}

type RecognitionBadge struct {
	BaseModelSeq
	Name        *string
	Description *string `gorm:"size:1024"`
}

type RecognitionBadgeStore struct {
	db *gorm.DB
}

func NewRecognitionStore(db *gorm.DB) *RecognitionStore {
	return &RecognitionStore{
		db: db,
	}
}

func (store *RecognitionStore) Save(recognition *Recognition) (err error) {
	if err = store.db.Create(recognition).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *RecognitionStore) Flush() (err error) {
	return store.db.Model(&Recognition{}).Delete(&Recognition{}).Error
}

func NewRecognitionBadgeStore(db *gorm.DB) *RecognitionBadgeStore {
	return &RecognitionBadgeStore{
		db: db,
	}
}

func (store *RecognitionBadgeStore) Save(badge *RecognitionBadge) (err error) {
	if err = store.db.Create(badge).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *RecognitionBadgeStore) Flush() (err error) {
	return store.db.Model(&RecognitionBadge{}).Delete(&RecognitionBadge{}).Error
}
