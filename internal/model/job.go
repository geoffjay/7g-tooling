package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Job struct {
	BaseModelSeq
	Name  *string
	Tasks []Task `gorm:"many2many:job_tasks;association_autoupdate:false;association_autocreate:false"`
}

type JobStore struct {
	db *gorm.DB
}

func NewJobStore(db *gorm.DB) *JobStore {
	return &JobStore{
		db: db,
	}
}

func (store *JobStore) Save(job *Job) (err error) {
	if err = store.db.Create(job).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *JobStore) Flush() (err error) {
	return store.db.Model(&Job{}).Delete(&Job{}).Error
}
