package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Department struct {
	BaseModelSeq
	Name *string
	Lead *string
}

type DepartmentStore struct {
	db *gorm.DB
}

func NewDepartmentStore(db *gorm.DB) *DepartmentStore {
	return &DepartmentStore{
		db: db,
	}
}

func (store *DepartmentStore) Save(department *Department) (err error) {
	if err = store.db.Create(department).Error; err != nil {
		logrus.Panic(err)
	}
	return
}
