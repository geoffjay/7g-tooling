package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Role struct {
	gorm.Model
	SgID                 int
	Title                string
	Description          string
	ResponsibilityTitles []string             `gorm:"-" mapstructure:"role-responsibilities"`
	Responsibilities     []RoleResponsibility `gorm:"foreignkey:RoleID"`
}

type RoleStore struct {
	db *gorm.DB
}

type RoleResponsibility struct {
	gorm.Model
	Title          string
	Description    string
	RoleID         int
	RoleTemplateID int
}

type RoleResponsibilityStore struct {
	db *gorm.DB
}

type RoleTemplate struct {
	gorm.Model
	SgID                 int
	Title                string
	Description          string
	ResponsibilityTitles []string             `gorm:"-" mapstructure:"role-responsibilities"`
	Responsibilities     []RoleResponsibility `gorm:"foreignkey:RoleTemplateID"`
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

func (store *RoleStore) Flush() (err error) {
	return store.db.Model(&Role{}).Delete(&Role{}).Error
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

func (store *RoleTemplateStore) Flush() (err error) {
	return store.db.Model(&RoleTemplate{}).Delete(&RoleTemplate{}).Error
}
