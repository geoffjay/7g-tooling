package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Location struct {
	gorm.Model
	SgID int
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

func (store *LocationStore) Flush() (err error) {
	return store.db.Model(&Location{}).Delete(&Location{}).Error
}

func (store *LocationStore) GetLocationByName(name string) (location Location, err error) {
	err = store.db.Model(&Location{}).Where("name = ?", name).First(&location).Error
	return
}
