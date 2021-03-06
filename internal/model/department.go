package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Department struct {
	gorm.Model
	SgID int
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

func (store *DepartmentStore) Flush() (err error) {
	return store.db.Model(&Department{}).Delete(&Department{}).Error
}

func (store *DepartmentStore) GetDepartmentByName(name string) (department Department, err error) {
	err = store.db.Model(&Department{}).Where("name = ?", name).First(&department).Error
	return
}
