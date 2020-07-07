package model

import (
	"database/sql/driver"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type TaskType int64

const (
	TASK_TYPE_NOT_A_VALUE TaskType = iota
	TASK_TYPE_CREATE_OBJECTIVE
	TASK_TYPE_UPDATE_OBJECTIVE
	TASK_TYPE_DELETE_OBJECTIVE
)

var types = [...]string{
	"Not a type",
	"Create objective",
	"Update objective",
	"Delete objective",
}

func (taskType TaskType) String() string {
	return types[taskType]
}

func (taskType *TaskType) Scan(value interface{}) error {
	*taskType = TaskType(value.(int64))
	return nil
}
func (taskType TaskType) Value() (driver.Value, error) { return int64(taskType), nil }

type Task struct {
	BaseModelSeq
	Type TaskType `gorm:"type:integer"`
}

type TaskStore struct {
	db *gorm.DB
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{
		db: db,
	}
}

func (store *TaskStore) Save(task *Task) (err error) {
	if err = store.db.Create(task).Error; err != nil {
		logrus.Panic(err)
	}
	return
}
