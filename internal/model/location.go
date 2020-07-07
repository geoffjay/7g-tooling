package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Location struct {
	BaseModelSeq
	Name string
}

type LocationStore struct {
	db *gorm.DB
}

func NewLocationStore(db *gorm.DB) *LocationStore {
	return &LocationStore{
		db: db,
	}
}

func (store *LocationStore) Save(location *Location) (err error) {
	if err = store.db.Create(location).Error; err != nil {
		logrus.Panic(err)
	}
	return
}
