package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type KeyResult struct {
	gorm.Model
	SgID            int
	Name            *string `gorm:"not null"`
	MeasurementType int     `mapstructure:"measurement-type"`
	Weight          int     `mapstructure:"weight"`
	TargetValue     int     `mapstructure:"target-value"`
	StartingValue   int     `mapstructure:"starting-value"`
	ObjectiveID     int
}

type KeyResultStore struct {
	db *gorm.DB
}

type Objective struct {
	gorm.Model
	SgID         int
	Name         *string `gorm:"not null"`
	Type         string
	Description  string `gorm:"size:1024"`
	DueDate      string `mapstructure:"due-date"`
	CreationDate string `mapstructure:"creation-date"`
	StartDate    string `mapstructure:"start-date"`
	Draft        bool
	Closed       bool
	Private      bool
	//Department
	//Labels []string
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
	KeyResults []*KeyResult `gorm:"foreignkey:ObjectiveID" mapstructure:"key-results"`

	// Ignored fields
	DepartmentSgIDs  []int `gorm:"-"`
	ParentSgIDs      []int `gorm:"-"`
	ChildSgIDs       []int `gorm:"-"`
	OwnerSgIDs       []int `gorm:"-"`
	StakeholderSgIDs []int `gorm:"-"`
	FollowerSgIDs    []int `gorm:"-"`
}

type ObjectiveStore struct {
	db      *gorm.DB
	KRStore *KeyResultStore
}

func NewKeyResultStore(db *gorm.DB) *KeyResultStore {
	return &KeyResultStore{
		db: db,
	}
}

func (store *KeyResultStore) Save(keyResult *KeyResult) (err error) {
	if err = store.db.Create(keyResult).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *KeyResultStore) Flush() (err error) {
	return store.db.Model(&KeyResult{}).Delete(&KeyResult{}).Error
}

func NewObjectiveStore(db *gorm.DB) *ObjectiveStore {
	return &ObjectiveStore{
		db:      db,
		KRStore: NewKeyResultStore(db),
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

func (o *Objective) GetTypeID() int {
	// TODO: return int for type name
	return 1
}
