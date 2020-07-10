package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type OneOnOne struct {
	gorm.Model
	SgID int
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
	gorm.Model
	SgID        int
	Name        *string    `gorm:"not null"`
	Description *string    `gorm:"size:1024"`
	Questions   []Question `gorm:"foreignkey:OneOnOneTemplateID"`
}

type OneOnOneTemplateStore struct {
	db *gorm.DB
}

type Question struct {
	gorm.Model
	SgID        int
	Content     *string
	Description *string `gorm:"size:1024"`
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
	for _, question := range template.Questions {
		store.db.Save(&question)
	}
	return
}

func (store *OneOnOneTemplateStore) Flush() (err error) {
	return store.db.Model(&OneOnOneTemplate{}).Delete(&OneOnOneTemplate{}).Error
}
