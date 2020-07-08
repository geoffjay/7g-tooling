package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type OneOnOne struct {
	BaseModelSeq
	//1-on-1 Template Name
	//1-on-1 Template ID
	//1-on-1 Creation Date
	//1-on-1 Scheduled Date
	//1-on-1 Completion Date
	Facilitator User
	TeamMember  User
	//Question 1
	//Facilitator Notes
	//Team Member Notes
	//Question 2
	//Facilitator Notes
	//Team Member Notes
	//Question 3
	//Facilitator Notes
	//Team Member Notes
	//Question 4
	//Facilitator Notes
	//Team Member Notes
	//Question 5
	//Facilitator Notes
	//Team Member Notes
}

type OneOnOneStore struct {
	db *gorm.DB
}

type OneOnOneTemplate struct {
	BaseModelSeq
	Name        *string `gorm:"not null"`
	Description *string `gorm:"size:1024"`
	//1on1 Frequency
	//Question Name
	//Question Description `gorm:"size:1024"`
}

type OneOnOneTemplateStore struct {
	db *gorm.DB
}

func NewOneOnOneStore(db *gorm.DB) *OneOnOneStore {
	return &OneOnOneStore{
		db: db,
	}
}

func (store *OneOnOneStore) Save(oneonone *OneOnOne) (err error) {
	if err = store.db.Create(oneonone).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *OneOnOneStore) Flush() (err error) {
	return store.db.Model(&OneOnOne{}).Delete(&OneOnOne{}).Error
}

func NewOneOnOneTemplateStore(db *gorm.DB) *OneOnOneTemplateStore {
	return &OneOnOneTemplateStore{
		db: db,
	}
}

func (store *OneOnOneTemplateStore) Save(template *OneOnOneTemplate) (err error) {
	if err = store.db.Create(template).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *OneOnOneTemplateStore) Flush() (err error) {
	return store.db.Model(&OneOnOneTemplate{}).Delete(&OneOnOneTemplate{}).Error
}
