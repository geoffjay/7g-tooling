package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Competency struct {
	BaseModelSeq
	Title       *string
	Description *string `gorm:"size:1024"`
	//Core Competency
	Levels []Level `gorm:"many2many:competency_levels;association_autoupdate:false;association_autocreate:false"`
}

type CompetencyStore struct {
	db *gorm.DB
}

type Level struct {
	BaseModelSeq
	Name        *string
	Description *string `gorm:"size:1024"`
}

type LevelStore struct {
	db *gorm.DB
}

func NewCompetencyStore(db *gorm.DB) *CompetencyStore {
	return &CompetencyStore{
		db: db,
	}
}

func (store *CompetencyStore) Save(competency *Competency) (err error) {
	if err = store.db.Create(competency).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func NewLevelStore(db *gorm.DB) *LevelStore {
	return &LevelStore{
		db: db,
	}
}

func (store *LevelStore) Save(level *Level) (err error) {
	if err = store.db.Create(level).Error; err != nil {
		logrus.Panic(err)
	}
	return
}
