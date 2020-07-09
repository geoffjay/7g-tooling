package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Objective struct {
	gorm.Model
	SgID        int
	Name        *string `gorm:"not null"`
	Type        *string
	Description *string `gorm:"size:1024"`
	//Due Date
	//Creation Date
	//Open/Closed
	//Privacy
	//Objective Department
	//Labels
	//Parent Objective ID(s)
	//Child Objective ID(s)
	//Child Objective Weight(s)
	//Creator name
	//Creator ID
	//Owner name(s)
	//Owner ID(s)
	//Stakeholder name(s)
	//Stakeholder ID(s)
	//Follower name(s)
	//Follower ID(s)
	//KR name(s)
	//KR Unit(s)
	//KR Weight(s)
	//KR Target Value(s)
	//KR Start Value(s)
}

type ObjectiveStore struct {
	db *gorm.DB
}

func NewObjectiveStore(db *gorm.DB) *ObjectiveStore {
	return &ObjectiveStore{
		db: db,
	}
}

func (store *ObjectiveStore) Save(objective *Objective) (err error) {
	if err = store.db.Create(objective).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *ObjectiveStore) Flush() (err error) {
	return store.db.Model(&Objective{}).Delete(&Objective{}).Error
}
