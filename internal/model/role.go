package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Role struct {
	BaseModelSeq
	Name string
}

type RoleStore struct {
	db *gorm.DB
}

type RoleTemplate struct {
	BaseModelSeq
	Name string
	//Description
	//Responsibility 1 Title
	//Responsibility 1 Description
	//Responsibility 2 Title
	//Responsibility 2 Description
	//Responsibility 3 Title
	//Responsibility 3 Description
	//Responsibility 4 Title
	//Responsibility 4 Description
	//Responsibility 5 Title
	//Responsibility 5 Description
	//Responsibility 6 Title
	//Responsibility 6 Description
	//Responsibility 7 Title
	//Responsibility 7 Description
	//Responsibility 8 Title
	//Responsibility 8 Description
	//Responsibility 9 Title
	//Responsibility 9 Description
	//Responsibility 10 Title
	//Responsibility 10 Description
	//Responsibility 11 Title
	//Responsibility 11 Description
}

type RoleTemplateStore struct {
	db *gorm.DB
}

func NewRoleStore(db *gorm.DB) *RoleStore {
	return &RoleStore{
		db: db,
	}
}

func (store *RoleStore) Save(role *Role) (err error) {
	if err = store.db.Create(role).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func NewRoleTemplateStore(db *gorm.DB) *RoleTemplateStore {
	return &RoleTemplateStore{
		db: db,
	}
}

func (store *RoleTemplateStore) Save(template *RoleTemplate) (err error) {
	if err = store.db.Create(template).Error; err != nil {
		logrus.Panic(err)
	}
	return
}
